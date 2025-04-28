package puzzle

import "testing"

func TestOptions(t *testing.T) {
	config := NewConfig()
	var x uint64 = 0
	DefineVar(config, "test", &x,
		WithEnvName("TEST"),
		WithDescription("a test option"),
		WithSliceSeparator("|"),
		WithFormat("???"),
		WithFlagName("test"),
		WithShortFlagName("t"),
		WithoutFlagName(),
	)
}
