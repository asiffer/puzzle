package puzzle

import (
	"fmt"
)

// A Converter is an object attached to a type. It may supports several frontends
// that are able to convert the entry to their own language
type Converter[T any] interface {
	Register(name Frontend, callback ConvertCallback[T]) error
	Convert(frontend Frontend, entry *Entry[T], args ...any) error
}

type ConvertCallback[T any] func(entry *Entry[T], args ...interface{}) error

func ConvertCallbackFactory1[T any, A any](fun func(entry *Entry[T], arg A) error) ConvertCallback[T] {
	return func(entry *Entry[T], args ...interface{}) error {
		var t T
		if len(args) < 1 {
			return fmt.Errorf("an argument of type %T is required", t)
		}
		if len(args) > 1 {
			return fmt.Errorf("too many arguments provided, only an argument of type %T is required", t)
		}
		switch a := args[0].(type) {
		case A:
			return fun(entry, a)
		default:
			return fmt.Errorf("expected an argument of type %T, got %v (%T)", t, a, a)
		}
	}
}

// converter is a base converter
type converter[T any] struct {
	frontends map[Frontend]ConvertCallback[T]
}

// newConverter creates a new converter with a default string frontend
func newConverter[T any](fun func(entry *Entry[T], arg string) error) Converter[T] {
	return &converter[T]{
		frontends: map[Frontend]ConvertCallback[T]{
			StringFrontend: ConvertCallbackFactory1(fun), // we need a string convert at lea
		},
	}
}

func (c *converter[T]) Register(name Frontend, callback ConvertCallback[T]) error {
	_, exists := c.frontends[name]
	if exists {
		return fmt.Errorf("converter %s already exists", name)
	}
	c.frontends[name] = callback
	return nil
}

func (c *converter[T]) Convert(name Frontend, entry *Entry[T], args ...any) error {
	callback, exists := c.frontends[name]
	if !exists {
		var t T
		return fmt.Errorf("frontend %s does not exist (type: %T)", name, t)
	}
	return callback(entry, args...)
}
