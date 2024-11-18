package main

import (
	"os"

	"github.com/goplugin/plugin-env/config"
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
	"github.com/goplugin/plugin-env/pkg/helm/mockserver"
	mockservercfg "github.com/goplugin/plugin-env/pkg/helm/mockserver-cfg"
)

func main() {
	// see REMOTE_RUN.md for tutorial
	e := environment.New(&environment.Config{
		NamespacePrefix: "zmytest",
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 1,
			"plugin": map[string]interface{}{
				"image": map[string]interface{}{
					"image":   os.Getenv(config.EnvVarCLImage),
					"version": os.Getenv(config.EnvVarCLTag),
				},
			},
		}))
	if err := e.Run(); err != nil {
		panic(err)
	}
}
