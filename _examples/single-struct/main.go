package main

import (
	"time"

	"github.com/evenboee/table"
)

type User struct {
	ID             string    `table:"ID"`
	Name           string    `table:"Name"`
	Age            int       `table:"Age"`
	HashedPassword string    `table:"-"` // ignored
	CreatedAt      time.Time `table:"Created At"`
}

func main() {
	user := User{
		ID:             "a",
		Name:           "Alice",
		Age:            20,
		HashedPassword: "123456",
		CreatedAt:      time.Now(),
	}

	// Create a custom stringer
	customSingleStructStringer := table.StringerUnicode.WithOpts(
		table.WithOmitHeaders(true),
		table.WithSepAsSep(table.SepRow, table.SepHeader),
		table.WithRightAlignNumeric(false),
	)

	// Save opts for reuse
	opts := []table.StringifyConfigOpt{
		table.WithStringer(customSingleStructStringer),
		table.WithTimeFormat("15:04:05 02-01-2006"),
	}

	s := table.SingleStruct(user, opts...)
	println(s)
}
