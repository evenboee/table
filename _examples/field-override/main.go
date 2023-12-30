package main

import (
	"strings"
	"time"

	"github.com/evenboee/table"
)

type User struct {
	ID             string    `table:"ID"`
	Name           string    `table:"Name"`
	Age            int       `table:"Age"`
	HashedPassword string    `table:"-"` // ignored
	CreatedAt      time.Time `table:"Created At"`

	// New
	Children []User `table:"Children"`
}

// implement generalizer.Tabular interface
func (u User) ToTable() map[string]string {
	children := "-"
	if len(u.Children) > 0 {
		c := make([]string, len(u.Children))
		for i, child := range u.Children {
			c[i] = child.ID
		}
		children = strings.Join(c, ", ")
	}

	return map[string]string{
		"Children":      children,                                        // Override
		"Year of Birth": time.Now().AddDate(-u.Age, 0, 0).Format("2006"), // New
	}
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

	users[1].Children = []User{users[0], users[1]}

	s := table.String(users)
	println(s)
}
