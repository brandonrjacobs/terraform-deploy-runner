package app

import (
	"deploy-runner/config"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"text/template"
	"unicode"
)

// HelpWriter is the interface for writing out custom help with env var information
type HelpWriter interface {
	Write(w io.Writer, cmd *cobra.Command, envVars []config.EnvVar) error
}

// NewHelpWriter returns the default HelpWriter implementation
func NewHelpWriter() HelpWriter {
	return &helpWriter{}
}

var helpTemplate = `{{with (or .Command.Long .Command.Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Command.Runnable .Command.HasSubCommands}}{{.Command.UsageString}}{{end}}

Environment Variables: {{range .EnvVars}}{{. | listEnvVar}}{{end}}
`

type helpTemplateData struct {
	Command *cobra.Command
	EnvVars []config.EnvVar
}

type helpWriter struct {
}

func (h *helpWriter) Write(w io.Writer, cmd *cobra.Command, envVars []config.EnvVar) error {
	t := template.New("help")
	t.Funcs(templateFuncs)
	template.Must(t.Parse(helpTemplate))
	return t.Execute(w, helpTemplateData{Command: cmd, EnvVars: envVars})
}

var templateFuncs = template.FuncMap{
	"trimTrailingWhitespaces": trimRightSpace,
	"listEnvVar": listEnvVar,
}

func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func listEnvVar(env config.EnvVar) string {
	return "\n" + env.Name + ": " + env.Description
}

