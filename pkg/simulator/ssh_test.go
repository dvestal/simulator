package simulator_test

import (
	"testing"

	"github.com/controlplaneio/simulator-standalone/pkg/simulator"
	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	t.Skip("Need to mock out terraform output")
	t.Parallel()
	simulator := simulator.NewSimulator(
		simulator.WithLogger(noopLogger),
		simulator.WithTfDir(fixture("noop-tf-dir")),
		simulator.WithScenariosDir(fixture("valid")),
		simulator.WithAttackTag("Latest"),
		simulator.WithTfVarsDir("test"))

	cfg, err := simulator.SSHConfig()

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}
