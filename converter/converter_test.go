package converter

import (
	"strconv"
	"testing"
	"time"
)

func requireEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func pointTo[T any](v T) *T {
	return &v
}

func Test__Convert(t *testing.T) {
	gen := New()

	nilString := "value is nil"
	gen.Nil = func() string {
		return nilString
	}
	gen.Time = TimeFormatter(time.RFC3339)
	gen.Int = func(i int64) string {
		return strconv.FormatInt(i, 10)
	}
	gen.UInt = func(i uint64) string {
		return strconv.FormatUint(i, 10)
	}
	gen.Duration = func(d time.Duration) string {
		return d.String()
	}
	gen.Bool = BoolFormatter("yes", "no")

	requireEqual(t, nilString, gen.Convert(nil))

	var nilInt *int = nil
	requireEqual(t, nilString, gen.Convert(nilInt))

	var intPtr *int = pointTo(42)
	requireEqual(t, "42", gen.Convert(intPtr))

	var str string = "hello"
	requireEqual(t, "hello", gen.Convert(str))

	var i8 int8 = 8
	requireEqual(t, "8", gen.Convert(i8))
	requireEqual(t, "8", gen.Convert(&i8))

	var i16 int16 = 16
	requireEqual(t, "16", gen.Convert(i16))
	requireEqual(t, "16", gen.Convert(&i16))

	var i32 int32 = 32
	requireEqual(t, "32", gen.Convert(i32))
	requireEqual(t, "32", gen.Convert(&i32))

	var i64 int64 = 64
	requireEqual(t, "64", gen.Convert(i64))
	requireEqual(t, "64", gen.Convert(&i64))

	var ui8 uint8 = 8
	requireEqual(t, "8", gen.Convert(ui8))
	requireEqual(t, "8", gen.Convert(&ui8))

	var ui16 uint16 = 16
	requireEqual(t, "16", gen.Convert(ui16))
	requireEqual(t, "16", gen.Convert(&ui16))

	var ui32 uint32 = 32
	requireEqual(t, "32", gen.Convert(ui32))
	requireEqual(t, "32", gen.Convert(&ui32))

	var ui64 uint64 = 64
	requireEqual(t, "64", gen.Convert(ui64))
	requireEqual(t, "64", gen.Convert(&ui64))

	var f32 float32 = 3.14
	requireEqual(t, "3.14", gen.Convert(f32))
	requireEqual(t, "3.14", gen.Convert(&f32))

	var f64 float64 = 6.28
	requireEqual(t, "6.28", gen.Convert(f64))
	requireEqual(t, "6.28", gen.Convert(&f64))

	var b bool = true
	requireEqual(t, "yes", gen.Convert(b))
	requireEqual(t, "yes", gen.Convert(&b))

	b = false
	requireEqual(t, "no", gen.Convert(b))
	requireEqual(t, "no", gen.Convert(&b))

	var d time.Duration = 5 * time.Second
	requireEqual(t, "5s", gen.Convert(d))
	requireEqual(t, "5s", gen.Convert(&d))

	var tme time.Time = time.Date(1234, 5, 6, 12, 34, 56, 0, time.UTC)
	requireEqual(t, "1234-05-06T12:34:56Z", gen.Convert(tme))
	requireEqual(t, "1234-05-06T12:34:56Z", gen.Convert(&tme))
}

func Test__Convert_ConverterFunc(t *testing.T) {
	gen := New()
	gen.ConvertFunc = func(v any) (string, bool) {
		switch v := v.(type) {
		case int:
			return "int: " + strconv.Itoa(v), true
		}
		return "", false
	}

	i := 42
	requireEqual(t, "int: 42", gen.Convert(i))
	requireEqual(t, "int: 42", gen.Convert(&i))

	// see that is does not affect other types
	s := "abc"
	requireEqual(t, "abc", gen.Convert(s))
}

func Test__Convert_Fallback(t *testing.T) {
	gen := New()
	gen.Fallback = func(v any) string {
		return "fallback"
	}
	gen.Nil = nil

	a := struct{}{}
	requireEqual(t, "fallback", gen.Convert(a))
	requireEqual(t, "fallback", gen.Convert(&a))

	requireEqual(t, "fallback", gen.Convert(nil))

	var nilInt *int = nil
	requireEqual(t, "fallback", gen.Convert(nilInt))
}

func Test__StringCuroffFormatter(t *testing.T) {
	gen := New()

	gen.String = StringCutoffFormatter(4)
	requireEqual(t, "abcd", gen.Convert("abcd"))
	requireEqual(t, "a...", gen.Convert("abcde"))

	gen.String = StringCutoffFormatter(3)
	requireEqual(t, "abc", gen.Convert("abcd"))
	requireEqual(t, "abc", gen.Convert("abcde"))
}
