package v3

import (
	"fmt"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	cfgv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
)

var (
	pluginVersion            = plugin.Version{Number: 1}
	supportedProjectVersions = []config.Version{cfgv3.Version}
)

// TODO mbrauer - eventually, plugin.Full
var _ plugin.Init = Plugin{}

type Plugin struct {
	initSubcommand
}

func (p Plugin) Name() string {
	return "reconciler-runtime.mamachanko.com"
}

func (p Plugin) Version() plugin.Version {
	return pluginVersion
}

func (p Plugin) SupportedProjectVersions() []config.Version {
	return supportedProjectVersions
}

func (p Plugin) GetInitSubcommand() plugin.InitSubcommand {
	return &p.initSubcommand
}

var _ plugin.InitSubcommand = initSubcommand{}

type initSubcommand struct {
}

func (p *initSubcommand) UpdateMetadata(cliMeta plugin.CLIMetadata, subcmdMeta *plugin.SubcommandMetadata) {
	subcmdMeta.Description = `Initialize a new project including the following files:
  - a "go.mod" with project dependencies
  - a "PROJECT" file that stores project configuration
  - a "Makefile" with several useful make targets for the project
  - several YAML files for project deployment under the "config" directory
  - a "main.go" file that creates the manager that will run the project controllers
`
	subcmdMeta.Examples = fmt.Sprintf(`  # Initialize a new project with your domain and name in copyright
  %[1]s init --plugins go/v3 --domain example.org --owner "Your name"

  # Initialize a new project defining a specific project version
  %[1]s init --plugins go/v3 --project-version 3
`, cliMeta.CommandName)
}

func (p initSubcommand) Scaffold(fs machinery.Filesystem) error {
	//TODO implement me
	panic("implement me")
}
