package concurrent_test

import (
	"testing"

	"github.com/goplugin/plugin-env/environment"
	"github.com/goplugin/plugin-env/pkg/helm/plugin"
	"github.com/stretchr/testify/require"
)

func TestConcurrentEnvs(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		e := environment.New(nil).
			AddHelm(plugin.New(0, nil))
		defer e.Shutdown()
		err := e.Run()
		require.NoError(t, err)
	})
	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		e := environment.New(nil).
			AddHelm(plugin.New(0, nil))
		defer e.Shutdown()
		err := e.Run()
		require.NoError(t, err)
		err = e.
			ModifyHelm("plugin-0", plugin.New(0, map[string]interface{}{
				"replicas": 2,
			})).Run()
		require.NoError(t, err)
	})
}
