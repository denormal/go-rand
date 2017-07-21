package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestEndian(t *testing.T) {
	if rand.IsBigEndian() {
		t.Fatal("Expected IsBigEndian() to default to false")
	} else if !rand.IsLittleEndian() {
		t.Fatal("Expected IsLittleEndian() to default to true")
	}

	rand.BigEndian()
	if !rand.IsBigEndian() {
		t.Fatal("Expected IsBigEndian() to report true")
	} else if rand.IsLittleEndian() {
		t.Fatal("Expected IsLittleEndian() to report false")
	}

	rand.LittleEndian()
	if rand.IsBigEndian() {
		t.Fatal("Expected IsBigEndian() to report false")
	} else if !rand.IsLittleEndian() {
		t.Fatal("Expected IsLittleEndian() to report true")
	}
}

func TestBytes(t *testing.T) {
	var _maximum = 8 // the minimum byte allocation

	for _j := 0; _j < MAXLOOPS; _j++ {
		for _i := 0; _i < MAXBITS; _i++ {
			var (
				_length      = 1 << uint(_i)
				_bytes, _err = rand.Bytes(_length)
			)

			if _err != nil {
				t.Fatalf("Unexpected error from rand.Bytes(): %s", _err)
			} else if len(_bytes) != _length {
				t.Fatalf(
					"Expected bytes of length %d, got %d",
					_length, len(_bytes),
				)
			}

			if _maximum < _length {
				_maximum = _length
			}

			if cap(_bytes) < _maximum {
				t.Fatalf(
					"Expected bytes of at least capacity %d, got %d",
					_maximum, cap(_bytes),
				)
			}

			// ensure we don't have a slice of zero bytes
			//		- technically, that's a valid random value, but it's
			//		  highly unusual, so may be a sign of an underlying problem
			var _s string
			for _, _b := range _bytes {
				if _b != byte(0) {
					goto DONE
				}
			}
			_s = "s"
			if len(_bytes) == 1 {
				_s = ""
			}
			t.Errorf(
				"Expected non-zero random bytes; %d zero byte%s found",
				len(_bytes), _s,
			)

		DONE:
			rand.ReturnBytes(_bytes)
		}
	}

	_bytes, _err := rand.Bytes(-1)
	if _err != nil {
		t.Fatalf("Unexpected error from rand.Bytes(-1): %s", _err)
	} else if len(_bytes) != 0 {
		_s := "s"
		if len(_bytes) == 1 {
			_s = ""
		}

		t.Fatalf(
			"Expected zero length return from rand.Bytes(-1), got %d byte%s",
			len(_bytes), _s,
		)
	}
}
