package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestString(t *testing.T) {
	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length       = 1 << uint(_i)
			_string, _err = rand.String(_length)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.String(): %s", _err)
		} else if len(_string) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_string),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_string)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_string),
			)
		}

		// ensure some runes are non-zero
		var _s string
		for _, _r := range []rune(_string) {
			if _r != rune(0) {
				goto DONE
			}
		}
		_s = "s"
		if len([]rune(_string)) == 1 {
			_s = ""
		}
		t.Errorf(
			"Expected non-zero random string; %d zero rune%s found",
			len([]rune(_string)), _s,
		)

	DONE:
	}
}

func TestStringIf(t *testing.T) {
	var _cmp = func(r rune) bool { return r >= rune('a') && r <= rune('z') }

	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length       = 1 << uint(_i)
			_string, _err = rand.StringIf(_length, _cmp)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.String(): %s", _err)
		} else if len(_string) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_string),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_string)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_string),
			)
		}

		// ensure the runes are within the requested range
		for _, _r := range []rune(_string) {
			if !_cmp(_r) {
				t.Fatalf(
					"Expected only runes between 'a' and 'z'; found %q",
					string(_r),
				)
			}
		}
	}
}

func TestStringFrom(t *testing.T) {
	var _alphabet = "🤖.-"

	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length       = 1 << uint(_i)
			_string, _err = rand.StringFrom(_length, _alphabet)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.String(): %s", _err)
		} else if len(_string) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_string),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_string)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_string),
			)
		}

		// ensure the runes are within the requested range
		for _, _r := range []rune(_string) {
			switch _r {
			case '🤖':
			case '.':
			case '-':
			default:
				t.Fatalf(
					"Expected only runes of ',' and '-'; found %q",
					string(_r),
				)
			}
		}
	}
}
