package rand_test

import (
	"encoding/base64"
	"testing"

	"github.com/denormal/go-rand"
)

func TestBase64(t *testing.T) {
	for _i := 0; _i < MAXCOUNT; _i++ {
		_base64, _err := rand.Base64(_i)
		if _err != nil {
			t.Fatalf("Unexpected error from rand.Bas64(): %s", _err)
		} else {
			_bytes, _err := base64.RawStdEncoding.DecodeString(_base64)
			if _err != nil {
				t.Fatalf(
					"Unexpected error from "+
						"base64.StdEncoding.DecodeString(): %s", _err,
				)
			} else if len(_bytes) != _i {
				_s := "s"
				if len(_bytes) == 1 {
					_s = ""
				}

				t.Fatalf(
					"Expected %d byte%s from rand.Base64(), got %d",
					len(_bytes), _s, _i,
				)
			}
		}
	}

	_base64, _err := rand.Base64(-1)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.Base64(-1): %s", _err)
	} else if len(_base64) != 0 {
		t.Fatalf(
			"Expected zero length return from rand.Base64(-1), got %d",
			len(_base64),
		)
	}
}
