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

// ConvertCallbackFactory1 is a function that turns a specific 1-argument function into a generic ConvertCallback
func ConvertCallbackFactory1[T any, A any](fun func(entry *Entry[T], arg A) error) ConvertCallback[T] {
	return func(entry *Entry[T], args ...interface{}) error {
		var t A
		if len(args) < 1 {
			return fmt.Errorf("an argument of type %T is required", t)
		}
		if len(args) > 1 {
			return fmt.Errorf("too many arguments provided, only an argument of type %T is required", t)
		}

		a, ok := args[0].(A)
		if ok {
			return fun(entry, A(a))
		}
		return fmt.Errorf("[ConvertCallbackFactory1] expected an argument of type %T, got %v (%T)", t, args[0], args[0])
	}
}

// ConvertCallbackFactory2 is a function that turns a specific 2-argument function into a generic ConvertCallback
func ConvertCallbackFactory2[T any, A any, B any](fun func(entry *Entry[T], arg0 A, arg1 B) error) ConvertCallback[T] {
	return func(entry *Entry[T], args ...interface{}) error {
		var t0 A
		var t1 B
		if len(args) < 2 {
			return fmt.Errorf("an argument of type %T and an argument of type %T are required", t0, t1)
		}
		if len(args) > 2 {
			return fmt.Errorf("too many arguments provided, only an argument of type %T and an argument of type %T are required", t0, t1)
		}

		if a, ok := args[0].(A); ok {
			if b, ok := args[1].(B); ok {
				return fun(entry, a, b)
			} else {
				return fmt.Errorf("expected an argument of type %T, got %v (%T)", t1, args[1], args[1])
			}
		} else {
			return fmt.Errorf("expected an argument of type %T, got %v (%T)", t0, args[0], args[0])
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

// Register registers a new frontend
func (c *converter[T]) Register(name Frontend, callback ConvertCallback[T]) error {
	_, exists := c.frontends[name]
	if exists {
		return fmt.Errorf("converter %s already exists", name)
	}
	c.frontends[name] = callback
	return nil
}

// Convert calls the specific frontend on a given entry
func (c *converter[T]) Convert(name Frontend, entry *Entry[T], args ...any) error {
	callback, exists := c.frontends[name]
	if !exists {
		var t T
		return fmt.Errorf("frontend %s does not exist (type: %T)", name, t)
	}
	return callback(entry, args...)
}
