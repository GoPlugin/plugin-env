package main

import (
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/examples/deployment_part_cdk8s"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(nil).
		AddChart(deployment_part_cdk8s.New(&deployment_part_cdk8s.Props{})).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 2,
		}))
	if err := e.Run(); err != nil {
		panic(err)
	}
	e.Shutdown()
}
