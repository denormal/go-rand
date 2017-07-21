package rand_test

import (
	"encoding/hex"
	"testing"

	"github.com/denormal/go-rand"
)

func TestHex(t *testing.T) {
	for _i := 0; _i < MAXCOUNT; _i++ {
		_hex, _err := rand.Hex(_i)
		if _err != nil {
			t.Fatalf("Unexpected error from rand.Bytes(): %s", _err)
		} else {
			_bytes, _err := hex.DecodeString(_hex)
			if _err != nil {
				t.Fatalf(
					"Unexpected error from hex.DecodeString(): %s", _err,
				)
			} else if len(_bytes) != _i {
				_s := "s"
				if len(_bytes) == 1 {
					_s = ""
				}

				t.Fatalf(
					"Expected %d byte%s from rand.Hex(), got %d",
					len(_bytes), _s, _i,
				)
			}
		}
	}

	_hex, _err := rand.Hex(-1)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.Hex(-1): %s", _err)
	} else if len(_hex) != 0 {
		t.Fatalf(
			"Expected zero length return from rand.Hex(-1), got %d",
			len(_hex),
		)
	}
}
