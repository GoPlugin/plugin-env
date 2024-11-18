package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(&environment.Config{
		Labels: []string{fmt.Sprintf("envType=%s", pkg.EnvTypeEVM5)},
	}).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, nil))
	err := e.Run()
	if err != nil {
		panic(err)
	}
	// default k8s selector
	summ, err := e.ResourcesSummary("app in (plugin-0, geth)")
	if err != nil {
		panic(err)
	}
	log.Warn().Interface("Resources", summ).Send()
	e.Shutdown()
}
