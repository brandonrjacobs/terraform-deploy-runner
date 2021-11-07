package app

import (
	"context"
	"deploy-runner/config"
	"deploy-runner/internal"
	"errors"
	"go.uber.org/fx"

	"time"
)

var startTimeout = time.Second * 30
var stopTimeout = time.Second * 30

// Service is an abstraction for a long running process within the app, could be an http server or a background cache
// refresh thread etc
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Disabled() bool
}

// ServiceCommand creates a new CLI sub command that runs a long running service/dameon type process that would be
// made up of multiple Service interface implementations
func ServiceCommand(use, short, long string, writer HelpWriter) Command {
	return newCommand(use, short, long, false, newBuilderForService(), writer, []config.EnvVar{})
}

// ServiceRunner is the implementation of Runner for a ServiceCommand
type ServiceRunner struct {
	fx.In
	App appController
	Log internal.StartupLog
}

func (r *ServiceRunner) Run(args []string) error {
	// TODO: handle logging/error wrapping in here
	startCtx, cancel := context.WithTimeout(context.Background(), startTimeout)
	defer cancel()

	if err := r.App.Start(startCtx); err != nil {
		r.Log.Err(err, "Error starting app")
		return errors.New("unable to start app")
	}

	r.App.WaitForSignal()

	stopCtx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()

	if err := r.App.Stop(stopCtx); err != nil {
		r.Log.Err(err, "Error stopping app")
		return errors.New("unable to stop app")
	}

	return nil
}

type serviceContainer struct {
	fx.In
	Services []Service `group:"services"`
}

func runServices(c serviceContainer, lc fx.Lifecycle) {
	for _, service := range c.Services {
		if service.Disabled(){
			continue
		}
		lc.Append(fx.Hook{
			OnStart: service.Start,
			OnStop:  service.Stop,
		})
	}
}

func validateServices(c serviceContainer) error {
	if len(c.Services) < 1 {
		return errors.New("no services have been registered, you must register at least one Service implementation")
	}

	return nil
}

