package services

import (
	"tracks-grpc/config"
	"tracks-grpc/internal"
)

var Component = internal.NewComponent("services", []config.EnvVar{}, NewServer)
