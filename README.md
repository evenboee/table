# Tabular

Print structs, maps and arrays as tables. Primary focus is on structs. 

Inspiered by [https://ozh.github.io/ascii-tables/](https://ozh.github.io/ascii-tables/)

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
    - **generalizer.StringerConfig**
        - defines how will the generalizer will turn values into strings
        - default uses RFC3339 time format and turns nil to empty strings
    - offset and limit (use StructLimit to print in groups with a size of the limit)


TODO:
- [ ] Order additional columns (alphabetically?)
- [ ] due for a rework
