package puzzle

import "fmt"

type KeyAlreadyExistsError struct {
	Key string
}

func (e *KeyAlreadyExistsError) Error() string {
	return fmt.Sprintf("key %s already exists", e.Key)
}

type KeyNotFoundError struct {
	Key string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("key %s not found", e.Key)
}

type TypeMismatchError struct {
	Key           string
	RequestedType string
	ActualType    string
}

func newTypeMismatchError(key string, requested any, actual any) *TypeMismatchError {
	return &TypeMismatchError{
		Key:           key,
		RequestedType: fmt.Sprintf("%T", requested),
		ActualType:    fmt.Sprintf("%T", actual),
	}
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch for key %s: requested %s, got %s", e.Key, e.RequestedType, e.ActualType)
}

type InvalidValueError struct {
	Key   string
	Value any
	Err   error
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("invalid value for key %s: %v, error: %v", e.Key, e.Value, e.Err)
}
