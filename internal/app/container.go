package app

import "deploy-runner/config"

func ContainerCommand(use string, short, long string, writer HelpWriter) Command {
	return newCommand(use, short, long, true, nil, writer, []config.EnvVar{})
}
