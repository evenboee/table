# Table

Generate tables. Either explicitly from headers and rows or automatically generate from any type. 

```go
table.Generate([]string{"Name", "Age"}, [][]string{{"John", "25"}, {"Jane", "23"}}).Print()
/*
+------+-----+
| Name | Age |
+------+-----+
| John |  25 |
| Jane |  23 |
+------+-----+
*/
```

For structs, maps, slices and others:

```go
table.String([]Person{
    {"Alice", 23},
    {"Bob", 25},
}).Print()
/*
+-------+-----+
| Name  | Age |
+-------+-----+
| Alice |  23 |
| Bob   |  25 |
+-------+-----+
*/
```

With unicode table format and spreadsheet mode

```go
table.StringWith(table.StyleUnicode, []Person{
    {"Alice", 23},
    {"Bob", 25},
}, table.WithSpreadsheet("#")).Print()
/*
╔═════╦═════════╦═══════╗
║  #  ║  Name   ║  Age  ║
╠═════╬═════════╬═══════╣
║  1  ║  Alice  ║   23  ║
║  2  ║  Bob    ║   25  ║
╚═════╩═════════╩═══════╝
*/
```

Using struct tags and options. 
Use the "table" tag (configurable) to control what value is used in the header for a field. 
Fields with tag value "-" will be skipped. 
Included and excluded fields can be explicitly specified with generalizer.Include and generalizer.Exclude or convenience functions table.Include and table.Exclude. 
The include and exclude options use the field name. 
If include is specified only fields in the list will be included, if exclude is specified those fields will not be included. 
Include is applied first and then exclude. 

```go
type User struct {
    ID           string
    DisplayName  string `table:"Display Name"`
    Password     string `table:"-"`
    CreatedAt    time.Time
}

table.String([]User{
    {"1", "John Doe", "johnspwd", time.Now()},
    {"2", "Jane Doe", "janespwd", time.Now()},
}, table.Exclude("DisplayName"), table.TimeFormat("2006-01-02 15:04:05")).Print()
/*
+----+---------------------+
| ID |      CreatedAt      |
+----+---------------------+
|  1 | 2020-01-02 20:30:09 |
|  2 | 2020-01-02 20:30:09 |
+----+---------------------+
*/
```

In the previous example "CreatedAt" was just the value of the header. 
Use the table.Params.HeaderFormatter to format header. table.AutoFormatHeader can be applied with `table.WithAutoFormatHeader()`

```go
table.String([]User{
    {"1", "John Doe", "johnspwd", time.Now()},
    {"2", "Jane Doe", "janespwd", time.Now()},
},
    table.Exclude("DisplayName"),
    table.TimeFormat("2006-01-02 15:04:05"),
    table.WithAutoFormatHeader(),
).Print()
/*
+----+---------------------+
| ID |     Created At      |
+----+---------------------+
|  1 | 2020-01-02 20:30:09 |
|  2 | 2020-01-02 20:30:09 |
+----+---------------------+
*/
```

To reuse same configuration:
```go
sc := table.New(
    table.Exclude("DisplayName"),
    table.TimeFormat("2006-01-02 15:04:05"),
    table.WithAutoFormatHeader(),
)

sc.String([]User{
    {"1", "John Doe", "johnspwd", time.Now()},
}).Print()

sc.String([]User{
    {"2", "Jane Doe", "janespwd", time.Now()},
}).Print()

/*
+----+---------------------+
| ID |     Created At      |
+----+---------------------+
|  1 | 2020-01-02 20:30:09 |
+----+---------------------+
+----+---------------------+
| ID |     Created At      |
+----+---------------------+
|  2 | 2020-01-02 20:30:09 |
+----+---------------------+
*/
```
