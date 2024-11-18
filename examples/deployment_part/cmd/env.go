package main

import (
	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/examples/deployment_part"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"time"
)

func main() {
	e := environment.New(&environment.Config{
		NamespacePrefix:   "adding-new-deployment-part",
		TTL:               3 * time.Hour,
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(deployment_part.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 5,
			"env": map[string]interface{}{
				"SOLANA_ENABLED":              "true",
				"EVM_ENABLED":                 "false",
				"EVM_RPC_ENABLED":             "false",
				"PLUGIN_DEV":               "false",
				"FEATURE_OFFCHAIN_REPORTING2": "true",
				"feature_offchain_reporting":  "false",
				"P2P_NETWORKING_STACK":        "V2",
				"P2PV2_LISTEN_ADDRESSES":      "0.0.0.0:6690",
				"P2PV2_DELTA_DIAL":            "5s",
				"P2PV2_DELTA_RECONCILE":       "5s",
				"p2p_listen_port":             "0",
			},
		}))
	if err := e.Run(); err != nil {
		panic(err)
	}
}
