package app

import (
	"deploy-runner/config"
	"deploy-runner/internal"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"sort"
	"time"
)


// Command represents CLI command that can be run. This abstracts the config for the command and ultimately
// builds a cobra.Command under the hood so we can leverage that framework.
type Command interface {
	// Builds a cobra.Command from this command
	ToCobra(cfg *viper.Viper) *cobra.Command

	// AddComponent will add one or more components to the command so its dependencies are available for injection
	AddComponent(component ...internal.Component)

	// Adds one or more commands as sub commands of this one
	AddCommand(command ...Command)

	// Sets the parent command of this command
	SetParent(parent Command)

	// Adds an integer flag to the command
	IntFlag(viperKey, flag, usage string)

	// Adds a string flag to the command
	StringFlag(viperKey, flag, usage string)

	// Adds a bool flag to the command
	BoolFlag(viperKey, flag, usage string)

	// Adds a duration flag to the command
	DurationFlag(viperKey, flag, usage string)

	// Validates the dependencies for the command
	Validate(cfg *viper.Viper) error

	// Binds single environment variable to command
	BindEnv(envVar config.EnvVar)
}


type command struct {
	container     bool
	cobra         *cobra.Command
	components    []internal.Component
	flagBindings  map[string]string
	childCommands []Command
	parent        *command
	builder       Builder
	help          HelpWriter
	provide       fx.Option
	extraEnv      []config.EnvVar
}


func newCommand(use, short, long string, container bool, builder Builder, help HelpWriter, extraEnv []config.EnvVar, constructors ...interface{}) Command {
	var option fx.Option
	if len(constructors) > 0 {
		option = fx.Provide(constructors...)
	}

	return &command{
		container: container,
		cobra: &cobra.Command{
			Use:   use,
			Short: short,
			Long:  long,
		},
		flagBindings:  make(map[string]string),
		childCommands: make([]Command, 0),
		components:    make([]internal.Component, 0),
		builder:       builder,
		help:          help,
		provide:       option,
		extraEnv:      extraEnv,
	}
}

func (c *command) ToCobra(cfg *viper.Viper) *cobra.Command {
	components := c.mergeComponents()
	envVars := c.mergeEnvVars(components)

	if c.container {
		for _, cmd := range c.childCommands {
			c.cobra.AddCommand(cmd.ToCobra(cfg))
		}
	} else {
		flagBindings := c.mergeFlagBindings()

		c.cobra.PreRunE = func(cmd *cobra.Command, args []string) error {
			for f, v := range flagBindings {
				flag := cmd.Flags().Lookup(f)
				if flag == nil {
					return fmt.Errorf("bad flag binding flag %v not found", f)
				}
				if err := cfg.BindPFlag(v, flag); err != nil {
					return fmt.Errorf("unable to bind flag %v to viper key %v: %v", f, v, err)
				}
			}

			for _, v := range envVars {
				if err := cfg.BindEnv(v.Key.String(), v.Name); err != nil {
					return fmt.Errorf("unable to bind env var %v to viper key %v: %v", v.Name, v.Key.String(), err)
				}
			}

			return nil
		}

		options := c.mergeOptions(components)
		if c.provide != nil {
			options = append(options, c.provide)
		}

		c.cobra.RunE = func(cmd *cobra.Command, args []string) error {
			if runner, err := c.builder.BuildRunner(cfg, options...); err == nil {
				return runner.Run(args)
			} else {
				return err
			}
		}
	}

	c.cobra.SetHelpFunc(func(cmd *cobra.Command, strings []string) {
		if err := c.help.Write(cmd.OutOrStderr(), cmd, envVars); err != nil {
			fmt.Println("Error writing help", err)
		}
	})

	return c.cobra
}

func (c *command) AddComponent(components ...internal.Component) {
	c.components = append(c.components, components...)
}

func (c *command) AddCommand(command ...Command) {
	for _, cmd := range command {
		c.childCommands = append(c.childCommands, cmd)
		cmd.SetParent(c)
	}
}

func (c *command) SetParent(parent Command) {
	c.parent = parent.(*command)
}

func (c *command) IntFlag(viperKey, flag, usage string) {
	if c.container {
		c.cobra.PersistentFlags().Int(flag, 0, usage)
	} else {
		c.cobra.Flags().Int(flag, 0, usage)
	}

	c.flagBindings[flag] = viperKey
}

func (c *command) StringFlag(viperKey, flag, usage string) {
	if c.container {
		c.cobra.PersistentFlags().String(flag, "", usage)
	} else {
		c.cobra.Flags().String(flag, "", usage)
	}

	c.flagBindings[flag] = viperKey
}

func (c *command) BoolFlag(viperKey, flag, usage string) {
	if c.container {
		c.cobra.PersistentFlags().Bool(flag, false, usage)
	} else {
		c.cobra.Flags().Bool(flag, false, usage)
	}

	c.flagBindings[flag] = viperKey
}

func (c *command) DurationFlag(viperKey, flag, usage string) {
	if c.container {
		c.cobra.PersistentFlags().Duration(flag, time.Second, usage)
	} else {
		c.cobra.Flags().Duration(flag, time.Second, usage)
	}

	c.flagBindings[flag] = viperKey

}

func (c *command) Validate(cfg *viper.Viper) error {
	components := c.mergeComponents()

	// Set some basic values to ensure logging is okay
	cfg.Set(config.LogLevel.String(), "info")
	cfg.Set(config.LogFormat.String(), "console")
	options := c.mergeOptions(components)
	if c.provide != nil {
		options = append(options, c.provide)
	}
	return c.builder.Validate(cfg, options...)
}

func (c *command) mergeFlagBindings() map[string]string {
	merged := make(map[string]string)
	current := c
	for current != nil {
		for k, v := range current.flagBindings {
			merged[k] = v
		}
		current = current.parent
	}
	return merged
}

func (c *command) mergeComponents() []internal.Component {
	// Dedupe first
	dedupe := make(map[string]internal.Component)

	current := c
	for current != nil {
		for _, comp := range current.components {
			dedupe[comp.Name()] = comp
		}
		current = current.parent
	}

	merged := make([]internal.Component, 0, len(dedupe))
	for _, v := range dedupe {
		merged = append(merged, v)
	}

	return merged
}

func (c *command) BindEnv(envVar config.EnvVar) {
	c.extraEnv = append(c.extraEnv, envVar)
}

func (c *command) mergeEnvVars(components []internal.Component) []config.EnvVar {
	// First dedupe variables
	vars := make(map[string]config.EnvVar)
	for _, comp := range components {
		for _, v := range comp.GetEnvVars() {
			vars[v.Name] = v
		}
	}

	for _, v := range c.extraEnv {
		vars[v.Name] = v
	}

	// Next get them in a sorted state
	s := make([]config.EnvVar, 0, len(vars))
	for _, v := range vars {
		s = append(s, v)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Name < s[j].Name
	})

	return s
}

func (c *command) mergeOptions(components []internal.Component) []fx.Option {
	merged := make([]fx.Option, 0, len(components))
	for _, comp := range components {
		merged = append(merged, comp.GetOption())
	}

	return merged
}

