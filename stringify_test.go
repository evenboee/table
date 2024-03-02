package table

import "testing"

func Test__Generate(t *testing.T) {
	s := New(
		WithSpreadsheet("#"),
	).Generate(
		[]string{"A", "B"},
		[][]string{
			{"12345", "2"},
			{"abø", "123"},
		},
	)

	expected := `+---+-------+-----+
| # |   A   |  B  |
+---+-------+-----+
| 1 | 12345 |   2 |
| 2 | abø   | 123 |
+---+-------+-----+`

	if s.String() != expected {
		t.Errorf("expected %s, got %s", expected, s.String())
	}
}
