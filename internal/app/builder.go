package app


import (
	"context"
	"deploy-runner/config"
	"deploy-runner/internal"
	"deploy-runner/internal/logging"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)


// Builder is used to build the fx.App and Runner for a command and also provides a way to validate the dependencies are
// all met.
type Builder interface {
	BuildRunner(cfg *viper.Viper, options ...fx.Option) (Runner, error)
	Validate(cfg *viper.Viper, options ...fx.Option) error
}

type builder struct {
	Runner           Runner
	RegisterServices bool
	DefaultOption    fx.Option
}

func newBuilderForAction() Builder {
	return &builder{
		Runner: &actionRunner{},
	}
}

func newBuilderForService() Builder {
	return &builder{
		Runner:           &ServiceRunner{},
		RegisterServices: true,
	}
}

func (b *builder) BuildRunner(cfg *viper.Viper, options ...fx.Option) (Runner, error) {
	log, err := logging.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize logging %w", err)
	}

	aw := &appWrapper{}
	app := fx.New(fx.Populate(b.Runner), fx.Options(options...), b.getDefaultOption(), b.getSuppliedOption(cfg, aw, log), b.getInvokes(), fx.Logger(newZapPrinter(log)))
	if err := app.Err(); err != nil {
		return nil, fmt.Errorf("unable to initialize fx app: %w", err)
	}
	aw.App = app

	return b.Runner, nil
}

func (b *builder) Validate(cfg *viper.Viper, options ...fx.Option) error {
	log, _ := logging.New(cfg)
	aw := &appWrapper{}
	app := fx.New(fx.Populate(b.Runner), fx.Options(options...), b.getDefaultOption(), b.getSuppliedOption(cfg, aw, log), b.getValidateInvokes(), fx.Logger(&nullPrinter{}))
	return app.Err()
}

func (b *builder) getSuppliedOption(cfg *viper.Viper, aw *appWrapper, log *zap.Logger) fx.Option {
	return fx.Options(fx.Provide(func() appController { return aw }), fx.Supply(cfg, log))
}

func (b *builder) getInvokes() fx.Option {
	if b.RegisterServices {
		return fx.Options(fx.Invoke(validateConfig, runServices))
	} else {
		return fx.Invoke(validateConfig)
	}
}

func (b *builder) getValidateInvokes() fx.Option {
	if b.RegisterServices {
		return fx.Options(fx.Invoke(validateValidators, validateServices))
	} else {
		return fx.Invoke(validateValidators)
	}
}

func (b *builder) getDefaultOption() fx.Option {
	if b.DefaultOption == nil {
		return fx.Provide(logging.NewStartUp)
	}

	return b.DefaultOption
}

type nullPrinter struct {
}

func (p *nullPrinter) Printf(template string, args ...interface{}) {

}

type zapPrinter struct {
	Log *zap.SugaredLogger
}

func (p *zapPrinter) Printf(template string, args ...interface{}) {
	p.Log.Debugf(template, args...)
}

func newZapPrinter(root *zap.Logger) fx.Printer {
	log := root.Named("fx").Sugar()
	return &zapPrinter{
		Log: log,
	}
}

type appController interface {
	Start(ctx context.Context) error
	WaitForSignal()
	Stop(ctx context.Context) error
}

type appWrapper struct {
	App *fx.App
}

func (w *appWrapper) Start(ctx context.Context) error {
	w.assertAppSet()

	return w.App.Start(ctx)
}

func (w *appWrapper) WaitForSignal() {
	w.assertAppSet()

	<-w.App.Done()
}

func (w *appWrapper) Stop(ctx context.Context) error {
	w.assertAppSet()

	return w.App.Stop(ctx)
}

func (w *appWrapper) assertAppSet() {
	if w.App == nil {
		panic("app not initialized")
	}
}

type validatorsContainer struct {
	fx.In
	Validators []config.Validator `group:"configValidators"`
}

func validateConfig(v validatorsContainer, log internal.StartupLog) error {
	passed := true
	for _, validator := range v.Validators {
		if err := validator.Validate(); err != nil {
			passed = false
			log.Err(err, "Config validator failed to validate")
		}
	}

	if !passed {
		return errors.New("one or more config.Validators failed to validate")
	}
	return nil
}

func validateValidators(v validatorsContainer) error {
	// We just want to see we can inject validators and there are no issues with dependencies for them
	return nil
}

