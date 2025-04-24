package puzzle

import (
	"testing"
)

var TRUE_VALUES = []string{"1", "t", "T", "TRUE", "true", "True"}
var FALSE_VALUES = []string{"0", "f", "F", "FALSE", "false", "False"}

func TestBoolConverter(t *testing.T) {
	for _, value := range TRUE_VALUES {
		if err := testConverter(BoolConverter, value, true); err != nil {
			t.Error(err)
		}
	}

	for _, value := range FALSE_VALUES {
		if err := testConverter(BoolConverter, value, false); err != nil {
			t.Error(err)
		}
	}
}
