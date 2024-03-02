package table

import (
	"fmt"

	"github.com/evenboee/table/generalizer"
)

type Table string

func (t Table) String() string {
	return string(t)
}

func (t Table) Print() {
	fmt.Println(t.String())
}

func (s *Params) String(v any) Table {
	return s.Generate(generalizer.Any(v))
}

func String(v any, opts ...ParamsOption) Table {
	return New(opts...).String(v)
}

func StringWith(s *TableStyle, v any, opts ...ParamsOption) Table {
	return New(append([]ParamsOption{WithStyle(s)}, opts...)...).String(v)
}
