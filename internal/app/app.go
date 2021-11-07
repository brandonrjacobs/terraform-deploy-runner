package app

import (
	"deploy-runner/config"
	"go.uber.org/fx"
)

// Runner is used to abstract the logic to run an fx.App based on if it is a simple singular action or a long running
// service
type Runner interface {
	Run(args []string) error
}


// ActionAdapter is an interface used to run an action from the CLI by gathering parameters from the environment or
// command line arguments.
type ActionAdapter interface {
	Execute(args []string) error
}


// ActionCommand creates a new CLI sub command that runs an action by using an ActionAdapter which is provided by
// adapterFunc which can be any fx compatible constructor/provide function meaning it can take in needed dependencies
func ActionCommand(use, short, long string, adapterFunc interface{}, envVars ...config.EnvVar) Command {
	return newCommand(use, short, long, false, newBuilderForAction(), NewHelpWriter(), envVars, adapterFunc)
}

type actionRunner struct {
	fx.In
	Adapter ActionAdapter
}

func (r *actionRunner) Run(args []string) error {
	return r.Adapter.Execute(args)
}
