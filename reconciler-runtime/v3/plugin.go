package v3

import (
	"github.com/mamachanko/tanzubuilder/reconciler-runtime/v3/templates"
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
	// TODO mbrauer - helpful help
}

func (p initSubcommand) PreScaffold(fs machinery.Filesystem) error {
	// TODO mbrauer - run pre-check on fs
	return nil
}

func (p initSubcommand) Scaffold(fs machinery.Filesystem) error {
	// TODO mbrauer - collect all errors

	if err := machinery.NewScaffold(fs).Execute(
		&templates.Makefile{},
		// TODO mbrauer - tools.go
		// TODO mbrauer - README.go
		// TODO mbrauer - ...
	); err != nil {
		return err
	}

	if err := machinery.NewScaffold(fs, machinery.WithFilePermissions(0100755)).Execute(
		&templates.MakefileHelp{},
	); err != nil {
		return err
	}

	return nil
}
