package generalizer

import "reflect"

type Result struct {
	Headers []string
	Rows    []map[string]string
}

type dereferencable[T any] interface {
	Kind() reflect.Kind
	Elem() T
}

// dereference is used for reflect types to get the underlying type without pointers
func dereference[T dereferencable[T]](d T) T {
	for d.Kind() == reflect.Pointer {
		d = d.Elem()
	}
	return d
}
