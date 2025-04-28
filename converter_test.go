package puzzle

import (
	"fmt"
	"io"
	"strconv"
	"testing"
)

type testType struct {
	data int64
}

func TestConvertCallbackFactory1(t *testing.T) {
	f := func(entry *Entry[string], arg string) error {
		return nil
	}
	callback := ConvertCallbackFactory1(f)

	e := NewEntry[string]("test")
	// good type
	if err := callback(e, "test"); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	// bad type
	if err := callback(e, 1); err == nil {
		t.Errorf("expected error, got nil")
	}

	// too many args
	if err := callback(e, "test", "test"); err == nil {
		t.Errorf("expected error, got nil")
	}
	// too few args
	if err := callback(e); err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestConvertCallbackFactory2(t *testing.T) {
	f := func(entry *Entry[string], arg string, arg1 string) error {
		return nil
	}
	callback := ConvertCallbackFactory2(f)

	e := NewEntry[string]("test")
	// good type
	if err := callback(e, "test", "test"); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	// bad type
	if err := callback(e, 1, 2); err == nil {
		t.Errorf("expected error, got nil")
	}

	// too many args
	if err := callback(e, "test", "test", "test"); err == nil {
		t.Errorf("expected error, got nil")
	}
	// too few args
	if err := callback(e, "test"); err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestRegister(t *testing.T) {
	testConverter := newConverter(func(entry *Entry[testType], stringValue string) error {
		i, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			return err
		}
		entry.ValueP.data = i
		return nil
	})

	printConverter := func(entry *Entry[testType], writer io.Writer) error {
		str := fmt.Sprintf("%+v\n", *entry.ValueP)
		_, err := writer.Write([]byte(str))
		return err
	}
	if err := testConverter.Register("print", ConvertCallbackFactory1(printConverter)); err != nil {
		t.Error(err)
	}
	// twice should fail
	if err := testConverter.Register("print", ConvertCallbackFactory1(printConverter)); err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestBadConverter(t *testing.T) {
	converter := StringConverter
	e := NewEntry[string]("test")
	if err := converter.Convert("unknown frontend", e, "data"); err == nil {
		t.Errorf("expected error, got nil")
	}
}
