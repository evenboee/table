package generalizer

import (
	"errors"
	"reflect"
)

var StructTableTag = "table"

// Used for overrides and aggregates
type Tabular interface {
	ToTable() map[string]string
}

var ErrTypeNotStruct = errors.New("type is not a struct")

func Struct[T any](data []T, s *StringerConfig) ([]string, []map[string]string) {
	t := dereference(reflect.TypeOf(*new(T)))

	if t.Kind() != reflect.Struct {
		panic(ErrTypeNotStruct)
	}

	// Intermediary map to keep track of headers and if they are numeric
	h := make(map[string]struct{})

	rows := make([]map[string]string, len(data))
	for i, row := range data {
		r := generalizeStruct(row, s)
		for k := range r {
			h[k] = struct{}{}
		}

		rows[i] = r
	}

	// Add all headers to slice
	// Makes sure they are in the correct order
	headers := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := getNameOfField(f)
		if name == "-" { // Skip fields that are ignored
			continue
		}

		if _, ok := h[name]; ok {
			headers = append(headers, name)
			delete(h, name)
		} else {
			// Even if the field was never used, it should still be added
			//   case for then len(data) == 0 should still show headers
			headers = append(headers, name)
		}
	}

	// Add any remaining headers
	for k := range h {
		headers = append(headers, k)
	}

	return headers, rows
}

func SingleStruct[T any](data T, s *StringerConfig) ([]string, []map[string]string) {
	headers := []string{"Field", "Value"}

	t := dereference(reflect.TypeOf(data))

	r := generalizeStruct(data, s)

	fields := make([]map[string]string, 0, len(r))
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := getNameOfField(f)
		if name == "-" { // Skip fields that are ignored
			continue
		}

		if _, ok := r[name]; ok {
			fields = append(fields, map[string]string{
				"Field": name,
				"Value": r[name],
			})
			delete(r, name)
		}
	}

	for k, v := range r {
		fields = append(fields, map[string]string{
			"Field": k,
			"Value": v,
		})
	}

	return headers, fields
}

// extract key value pairs of the fields of struct
func generalizeStruct[T any](data T, s *StringerConfig) map[string]string {
	t := dereference(reflect.TypeOf(data))
	val := dereference(reflect.ValueOf(data))

	r := make(map[string]string)

	// If the struct is a Tabular, use the ToTable method
	if tabular, ok := any(data).(Tabular); ok {
		r = tabular.ToTable()
		if r == nil { // If ToTable returns nil, make sure it is not nil
			r = make(map[string]string)
		}
	}

	// Loop through all fields of the struct
	for j := 0; j < t.NumField(); j++ {
		f := t.Field(j)
		name := getNameOfField(f)
		if name == "-" { // Skip fields that are ignored
			continue
		}

		// If field was not overridden by ToTable, add it to row map and headers map
		if _, ok := r[name]; !ok {
			r[name] = s.ToString(val.Field(j).Interface()) // fmt.Sprintf("%v", val.Field(j))
		}
	}

	return r
}

func getNameOfField(f reflect.StructField) string {
	name := f.Name
	if tag := f.Tag.Get(StructTableTag); tag != "" {
		name = tag
	}

	return name
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
