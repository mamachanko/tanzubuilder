package main

import (
	rrv3 "github.com/mamachanko/tanzubuilder/reconciler-runtime/v3"
	"log"
	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
)

func main() {
	// From https://github.com/kubernetes-sigs/kubebuilder/blob/master/cmd/main.go

	// Bundle plugin which built the golang projects scaffold by Kubebuilder go/v3
	//gov3Bundle, _ := plugin.NewBundle(
	//	golang.DefaultNameQualifier,
	//	plugin.Version{Number: 3},
	//	kustomizecommonv1.Plugin{},
	//	golangv3.Plugin{},
	//)

	rrv3Plugin := rrv3.Plugin{}

	c, err := cli.New(
		cli.WithCommandName("tanzubuilder"),
		cli.WithVersion("0.1.0"),
		cli.WithDescription(`CLI tool for building Kubernetes extensions and tools for the Tanzu ecosystem
`),
		// Register the plugins options which can be used to do the scaffolds via your CLI tool. See that we are using as example here the plugins which are implemented and provided by Kubebuilder
		cli.WithPlugins(
			rrv3Plugin,
		),
		// Defines what will be the default plugin used by your binary. It means that will be the plugin used if no info be provided such as when the user runs `kubebuilder init`
		cli.WithDefaultPlugins(cfgv3.Version, rrv3Plugin),
		// Define the default project configuration version which will be used by the CLI when none is informed by --project-version flag.
		cli.WithDefaultProjectVersion(cfgv3.Version),
		// Adds the completion option for your CLI
		cli.WithCompletion(),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
