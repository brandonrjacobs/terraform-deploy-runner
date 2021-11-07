package terraform

import (
	"github.com/hashicorp/terraform-exec/tfexec"
)

type client struct {
	tfClient *tfexec.Terraform
}

func NewClient(workDir, execPath string) *client {
	cl, err := tfexec.NewTerraform(workDir, execPath)
	if err != nil {

	}
	return &client{tfClient: cl}
}
