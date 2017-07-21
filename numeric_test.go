package rand_test

import (
	"testing"
)

func numeric(t *testing.T, name string, fn func() (interface{}, error)) {
	numericif(t, name, fn, nil)
}

func numericif(
	t *testing.T,
	name string,
	fn func() (interface{}, error),
	cmp func(interface{}) error,
) {
	_zero := 0
	for _i := 0; _i < MAXCOUNT; _i++ {
		_value, _err := fn()
		if _err != nil {
			t.Fatalf("Unexpected error from rand.%s(): %s", name, _err)
		} else {
			_iszero := false
			switch _v := _value.(type) {
			case int:
				_iszero = _v == 0
			case int8:
				_iszero = _v == 0
			case int16:
				_iszero = _v == 0
			case int32:
				_iszero = _v == 0
			case int64:
				_iszero = _v == 0
			case uint:
				_iszero = _v == 0
			case uint8:
				_iszero = _v == 0
			case uint16:
				_iszero = _v == 0
			case uint32:
				_iszero = _v == 0
			case uint64:
				_iszero = _v == 0
			case float32:
				_iszero = _v == 0
			case float64:
				_iszero = _v == 0
			case complex64:
				_iszero = real(_v) == 0 && imag(_v) == 0
			case complex128:
				_iszero = real(_v) == 0 && imag(_v) == 0
			}

			if _iszero {
				_zero++
			}

			if cmp != nil {
				_err = cmp(_value)
				if _err != nil {
					t.Fatalf(
						"Numeric error from cmp(): %s", _err,
					)
				}
			}
		}
	}

	if _zero == MAXCOUNT {
		t.Fatalf(
			"Only zero results from rand.%s() after %d attempts",
			name, MAXCOUNT,
		)
	}
}
