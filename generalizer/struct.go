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

func (c *Converter) Struct(data any) Result {
	v := dereference(reflect.ValueOf(data))
	t := v.Type()

	// make sure value if a slice of structs
	if t.Kind() != reflect.Slice && t.Elem().Kind() != reflect.Struct {
		panic(ErrTypeNotStruct)
	}

	t = t.Elem() // Get the type of the struct

	// Intermediary map to keep track of headers and if they are numeric
	h := make(map[string]struct{})

	d := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		d[i] = v.Index(i).Interface()
	}

	rows := make([]map[string]string, len(d))
	for i, row := range d {
		r := c.generalizeStruct(row)
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

	return Result{
		Headers: headers,
		Rows:    rows,
	}
}

func (c *Converter) SingleStruct(data any) Result {
	headers := []string{"Field", "Value"}

	t := dereference(reflect.TypeOf(data))

	r := c.generalizeStruct(data)

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

	return Result{
		Headers: headers,
		Rows:    fields,
	}
}

// extract key value pairs of the fields of struct
func (c *Converter) generalizeStruct(data any) map[string]string {
	val := dereference(reflect.ValueOf(data))
	t := val.Type()

	r := make(map[string]string)

	// If the struct is a Tabular, use the ToTable method
	if tabular, ok := data.(Tabular); ok {
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
			r[name] = c.ToString(val.Field(j).Interface()) // fmt.Sprintf("%v", val.Field(j))
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
