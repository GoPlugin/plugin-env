package plugin

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/goplugin/plugin-env/client"
	"github.com/goplugin/plugin-env/config"
	"github.com/goplugin/plugin-env/environment"
)

const (
	AppName              = "plugin"
	NodesLocalURLsKey    = "plugin_local"
	NodesInternalURLsKey = "plugin_internal"
	DBsLocalURLsKey      = "plugin_db"
)

type Props struct{}

type Chart struct {
	Name    string
	Index   int
	Path    string
	Version string
	Props   *Props
	Values  *map[string]any
}

func (m Chart) IsDeploymentNeeded() bool {
	return true
}

func (m Chart) GetName() string {
	return m.Name
}

func (m Chart) GetPath() string {
	return m.Path
}

func (m Chart) GetVersion() string {
	return m.Version
}

func (m Chart) GetProps() any {
	return m.Props
}

func (m Chart) GetValues() *map[string]any {
	return m.Values
}

func (m Chart) ExportData(e *environment.Environment) error {
	// fetching all apps with label app=plugin-${deploymentIndex}:${instanceIndex}
	pods, err := e.Fwd.Client.ListPods(e.Cfg.Namespace, fmt.Sprintf("app=%s", m.Name))
	if err != nil {
		return err
	}
	for i := 0; i < len(pods.Items); i++ {
		localConnection, err := e.Fwd.FindPort(fmt.Sprintf("%s:%d", m.Name, i), "node", "access").
			As(client.LocalConnection, client.HTTP)
		if err != nil {
			return err
		}
		e.URLs[NodesLocalURLsKey] = append(e.URLs[NodesLocalURLsKey], localConnection)
		log.Info().Str("Deployment", m.Name).Int("Node", i).Str("URL", localConnection).Msg("Local connection")

		remoteConnection, err := e.Fwd.FindPort(fmt.Sprintf("%s:%d", m.Name, i), "node", "access").
			As(client.RemoteConnection, client.HTTP)
		if err != nil {
			return err
		}
		e.URLs[NodesInternalURLsKey] = append(e.URLs[NodesInternalURLsKey], remoteConnection)
		if e.Cfg.InsideK8s {
			e.URLs[NodesLocalURLsKey] = e.URLs[NodesInternalURLsKey]
		}
		log.Info().Str("Deployment", m.Name).Int("Node", i).Str("URL", remoteConnection).Msg("Remote (in cluster) connection")

		dbLocalConnection, err := e.Fwd.FindPort(fmt.Sprintf("%s:%d", m.Name, i), "plugin-db", "postgres").
			As(client.LocalConnection, client.HTTP)
		if err != nil {
			return err
		}
		e.URLs[DBsLocalURLsKey] = append(e.URLs[DBsLocalURLsKey], dbLocalConnection)
		log.Info().Str("Deployment", m.Name).Int("Node", i).Str("URL", dbLocalConnection).Msg("DB local Connection")
	}
	return nil
}

func defaultProps() map[string]any {
	internalRepo := os.Getenv(config.EnvVarInternalDockerRepo)
	pluginImage := "public.ecr.aws/plugin/plugin"
	postgresImage := "postgres"
	if internalRepo != "" {
		pluginImage = fmt.Sprintf("%s/plugin", internalRepo)
		postgresImage = fmt.Sprintf("%s/postgres", internalRepo)
	}
	env := map[string]any{
		"CL_DATABASE_URL": "postgresql://postgres:verylongdatabasepassword@0.0.0.0/plugin?sslmode=disable",
	}
	pyroscopeAuth := os.Getenv(config.EnvVarPyroscopeKey)
	if pyroscopeAuth != "" {
		env["CL_PYROSCOPE_AUTH_TOKEN"] = pyroscopeAuth
	}
	return map[string]any{
		"replicas": "1",
		"env":      env,
		"plugin": map[string]any{
			"image": map[string]any{
				"image":   pluginImage,
				"version": "develop",
			},
			"web_port": "6688",
			"p2p_port": "8090",
			"resources": map[string]any{
				"requests": map[string]any{
					"cpu":    "350m",
					"memory": "1024Mi",
				},
				"limits": map[string]any{
					"cpu":    "350m",
					"memory": "1024Mi",
				},
			},
		},
		"db": map[string]any{
			"image": map[string]any{
				"image":   postgresImage,
				"version": "11.15",
			},
			"stateful": false,
			"capacity": "1Gi",
			"resources": map[string]any{
				"requests": map[string]any{
					"cpu":    "250m",
					"memory": "256Mi",
				},
				"limits": map[string]any{
					"cpu":    "250m",
					"memory": "256Mi",
				},
			},
		},
	}
}

func New(index int, props map[string]any) environment.ConnectedChart {
	return NewVersioned(index, "", props)
}

// NewVersioned enables you to select a specific helm chart version
func NewVersioned(index int, helmVersion string, props map[string]any) environment.ConnectedChart {
	dp := defaultProps()
	config.MustEnvOverrideVersion(&dp)
	config.MustMerge(&dp, props)
	return Chart{
		Index:   index,
		Name:    fmt.Sprintf("%s-%d", AppName, index),
		Path:    "plugin-qa/plugin",
		Version: helmVersion,
		Values:  &dp,
	}
}
