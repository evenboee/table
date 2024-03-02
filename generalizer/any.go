package generalizer

import "reflect"

func Any(v any, opts ...ParamsOption) ([]string, [][]string) {
	return New(opts...).Any(v)
}

func (params *Params) Any(v any) ([]string, [][]string) {
	val := reflect.ValueOf(v)
	return params.any_(val)
}

func (params *Params) any_(v reflect.Value) ([]string, [][]string) {
	if !v.IsValid() {
		return unknown(params.Converter.Nil(), "nil")
	}

	switch v.Kind() {
	case reflect.Struct:
		return params.SingleStruct(v.Interface())
	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Struct:
			return params.struct_(v)
		case reflect.Map:
			return params.map_(v)
		default:
			return params.slice_(v)
		}
	case reflect.Map:
		return params.singleMap_(v)
	case reflect.Pointer:
		if !v.IsNil() {
			return params.any_(v.Elem())
		}
	}

	return unknown(params.Converter.Convert(v.Interface()), v.Type().String())
}

func unknown(v string, t string) ([]string, [][]string) {
	return []string{"Key", "Value"},
		[][]string{
			{"Value", v},
			{"Type", t},
		}
}
