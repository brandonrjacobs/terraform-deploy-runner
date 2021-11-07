package services

import (
	"deploy-runner/config"
	"deploy-runner/internal"
	"deploy-runner/internal/app"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type serviceOut struct {
	fx.Out
	Service app.Service `group:"services"`
}

func NewServer(cfg *viper.Viper, log internal.BackgroundLog, rt *chi.Mux) serviceOut {
	addr := fmt.Sprintf(":%s", cfg.GetString(config.GrpcAddress.String()))
	return serviceOut{Service: &server{
		log:     log,
		address: addr,
		router:  rt,
	}}
}
