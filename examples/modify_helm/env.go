package main

import (
	"fmt"
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
	"github.com/goplugin/plugin-env/pkg/helm/mockserver"
	mockservercfg "github.com/goplugin/plugin-env/pkg/helm/mockserver-cfg"
)

func main() {
	e := environment.New(&environment.Config{
		NamespacePrefix: "modified-env",
		Labels:          []string{fmt.Sprintf("envType=Modified")},
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 1,
		}))
	err := e.Run()
	if err != nil {
		panic(err)
	}
	e.Cfg.KeepConnection = true
	e.Cfg.RemoveOnInterrupt = true
	err = e.
		ModifyHelm("plugin-0", plugin.New(0, map[string]interface{}{
			"replicas": 2,
		})).Run()
	if err != nil {
		panic(err)
	}
}
