# Tabular

Print structs, maps and arrays as tables. Primary focus is on structs. 

Inspired by [https://ozh.github.io/ascii-tables/](https://ozh.github.io/ascii-tables/)

See `_examples` for usage

In general:
- To override or add column to table: implement **generalizer.Tabular**
- the generalizer will turn a datastructure into:
    - list of headers to decide the order of the columns
    - list of rows consisting of a map of the columns
- For each stringify action is base on a **StringifyConfig** which includes:
    - **stringer** - defines the look of the table (default: DefaultStringer)
        - padding (makes all rows of column same width)
        - margin (distance from column content + padding to sides)
        - separators (top, row and bottom)
        - more...
    - **generalizer.Converter**
        - defines how will the generalizer will turn values into strings
        - default uses RFC3339 time format and turns nil to empty strings
    - offset and limit


## Basic Example

```go
type User struct {
	ID             string    `table:"ID"`
	Name           string    `table:"Name"`
	Age            int       `table:"Age"`
	HashedPassword string    `table:"-"` // ignored
	CreatedAt      time.Time `table:"Created At"`
}

users := []User{...}
println(table.String(users))
```

```
| ID | Name  | Age | Created At                |
+----+-------+-----+---------------------------+
| a  | Alice |  20 | 2023-11-16T23:02:38+01:00 |
| b  | Bob   |  21 | 2023-11-16T23:02:38+01:00 |
```


TODO:
- [ ] Order additional columns (alphabetically?)
- [ ] Table tag as option (specify other such as json, xml, yaml to use)
- [ ] consider rework
  - StringifyConfig is a little clunky to understand
  - Generalizer Converter configuration when using table.String is excessively verbose
- plan (and fix) undefined nil behavior
