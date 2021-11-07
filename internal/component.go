package internal

import (
	"deploy-runner/config"
	"go.uber.org/fx"
)

// Component is an abstraction of a group dependencies generally interfaces made up of their constructors used
// by fx for dependency injection and the env vars needed to configure them
type Component interface {
	Name() string
	GetOption() fx.Option
	GetEnvVars() []config.EnvVar
}

// NewComponent creates a new component made up of a set of environment variables and constructor functions to be
// used by FX to make the interfaces etc of the component available for dependency injection
func NewComponent(name string, envVars []config.EnvVar, constructors ...interface{}) Component {
	return &component{
		name:    name,
		option:  fx.Provide(constructors...),
		envVars: envVars,
	}
}

type component struct {
	name    string
	option  fx.Option
	envVars []config.EnvVar
}

func (c *component) Name() string {
	return c.name
}

func (c *component) GetOption() fx.Option {
	return c.option
}

func (c *component) GetEnvVars() []config.EnvVar {
	return c.envVars
}


