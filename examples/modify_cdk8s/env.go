package main

import (
	"fmt"
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg/cdk8s/blockscout"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(&environment.Config{
		NamespacePrefix: "modified-env",
		Labels:          []string{fmt.Sprintf("envType=Modified")},
	}).
		AddChart(blockscout.New(&blockscout.Props{
			WsURL:   "ws://geth:8546",
			HttpURL: "http://geth:8544",
		})).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 1,
		}))
	err := e.Run()
	if err != nil {
		panic(err)
	}
	e.ClearCharts()
	err = e.
		AddChart(blockscout.New(&blockscout.Props{
			HttpURL: "http://geth:9000",
		})).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 1,
		})).
		Run()
	if err != nil {
		panic(err)
	}
}
