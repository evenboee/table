package main

import (
	"time"

	table "github.com/evenboee/go-tabular"
)

type User struct {
	ID             string    `table:"ID"`
	Name           string    `table:"Name"`
	Age            int       `table:"Age"`
	HashedPassword string    `table:"-"` // ignored
	CreatedAt      time.Time `table:"Created At"`
}

func main() {
	users := []User{
		{
			ID:             "a",
			Name:           "Alice",
			Age:            20,
			HashedPassword: "123456",
			CreatedAt:      time.Now(),
		},
		{
			ID:             "b",
			Name:           "Bob",
			Age:            21,
			HashedPassword: "password",
			CreatedAt:      time.Now(),
		},
	}

	// Basic
	s := table.Struct(users)
	println(s)

	// Different stringer
	s = table.Struct(users, table.WithStringer(table.StringerUnicode))
	println(s)

	// Different time format
	startT := time.Now()
	s = table.Struct(users, table.WithTimeFormat("15:04:05 02-01-2006"))
	elapsed := time.Since(startT)
	println(s, elapsed)

	// Spreadsheet mode (row number)
	s = table.Struct(users, table.WithDefaultSpreadsheet) // or table.WithSpreadsheet("Row")
	println(s)
}
