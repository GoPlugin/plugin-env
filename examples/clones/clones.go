package main

import (
	"fmt"

	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
)

func main() {
	// Multiple environments of the same type/chart
	err := environment.New(&environment.Config{
		Labels:            []string{fmt.Sprintf("envType=%s", pkg.EnvTypeEVM5)},
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"plugin": map[string]interface{}{
				"resources": map[string]interface{}{
					"requests": map[string]interface{}{
						"cpu": "344m",
					},
					"limits": map[string]interface{}{
						"cpu": "344m",
					},
				},
			},
			"db": map[string]interface{}{
				"stateful": "true",
				"capacity": "1Gi",
			},
		})).
		AddHelm(plugin.New(1,
			map[string]interface{}{
				"plugin": map[string]interface{}{
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu": "577m",
						},
						"limits": map[string]interface{}{
							"cpu": "577m",
						},
					},
				},
			})).
		Run()
	if err != nil {
		panic(err)
	}
}
