package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestRunes(t *testing.T) {
	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length      = 1 << uint(_i)
			_runes, _err = rand.Runes(_length)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.Runes(): %s", _err)
		} else if len(_runes) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_runes),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_runes)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_runes),
			)
		}

		// ensure some runes are non-zero
		//		- only do this if we have sufficient runes
		if len(_runes) > 1 {
			var _s string
			for _, _r := range []rune(_runes) {
				if _r != rune(0) {
					goto DONE
				}
			}
			_s = "s"
			if len([]rune(_runes)) == 1 {
				_s = ""
			}
			t.Errorf(
				"Expected non-zero random string; %d zero rune%s found",
				len([]rune(_runes)), _s,
			)

		DONE:
		}
	}

	_runes, _err := rand.Runes(-1)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.Runes(-1): %s", _err)
	} else if len(_runes) != 0 {
		_s := "s"
		if len(_runes) == 1 {
			_s = ""
		}

		t.Fatalf(
			"Expected zero length return from rand.Runes(-1), got %d rune%s",
			len(_runes), _s,
		)
	}
}

func TestRunesIf(t *testing.T) {
	var _cmp = func(r rune) bool { return r >= rune('0') && r <= rune('9') }

	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length      = 1 << uint(_i)
			_runes, _err = rand.RunesIf(_length, _cmp)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.RunesIf(): %s", _err)
		} else if len(_runes) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_runes),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_runes)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_runes),
			)
		}

		// ensure the runes are within the requested range
		for _, _r := range []rune(_runes) {
			if !_cmp(_r) {
				t.Fatalf(
					"Expected only runes between '0' and '9'; found %q",
					string(_r),
				)
			}
		}
	}

	_runes, _err := rand.RunesIf(-1, _cmp)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.RunesIf(-1,...): %s", _err)
	} else if len(_runes) != 0 {
		_s := "s"
		if len(_runes) == 1 {
			_s = ""
		}

		t.Fatalf(
			"Expected zero length from rand.RunesIf(-1,...), got %d rune%s",
			len(_runes), _s,
		)
	}
}

func TestRunesFrom(t *testing.T) {
	var _alphabet = []rune("🤖.-")

	for _i := 0; _i < MAXBITS; _i++ {
		var (
			_length      = 1 << uint(_i)
			_runes, _err = rand.RunesFrom(_length, _alphabet)
		)

		if _err != nil {
			t.Fatalf("Unexpected error from rand.RunesFrom(): %s", _err)
		} else if len(_runes) < _length {
			t.Fatalf(
				"Expected string of at least length %d, got %s",
				_length, len(_runes),
			)
		}

		// ensure we have the requested number of runes
		if len([]rune(_runes)) != _length {
			_s := "s"
			if _length == 1 {
				_s = ""
			}

			t.Fatalf(
				"Expected string of %d rune%s, got %s",
				_length, _s, len(_runes),
			)
		}

		// ensure the runes are within the requested range
		for _, _r := range []rune(_runes) {
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

	_runes, _err := rand.RunesFrom(-1, _alphabet)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.RunesFrom(-1,...): %s", _err)
	} else if len(_runes) != 0 {
		_s := "s"
		if len(_runes) == 1 {
			_s = ""
		}

		t.Fatalf(
			"Expected zero length from rand.RunesFrom(-1,...), got %d rune%s",
			len(_runes), _s,
		)
	}
}
