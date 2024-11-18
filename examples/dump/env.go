package main

import (
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/goplugin/plugin-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(nil).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, nil))
	if err := e.Run(); err != nil {
		panic(err)
	}
	if err := e.DumpLogs("logs/mytest"); err != nil {
		panic(err)
	}
}
