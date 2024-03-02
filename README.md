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
