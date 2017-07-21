package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestBool(t *testing.T) {
	var (
		_true  int
		_false int
	)

	for _i := 0; _i < MAXCOUNT; _i++ {
		_bool, _err := rand.Bool()
		if _err != nil {
			t.Fatalf("Unexpected error from rand.Bool(): %s", _err)
		} else if _bool {
			_true++
		} else {
			_false++
		}
	}

	if _true == MAXCOUNT {
		t.Fatalf(
			"Expected true & false values, got %d true values only",
			_true,
		)
	} else if _false == MAXCOUNT {
		t.Fatalf(
			"Expected true & false values, got %d false values only",
			_false,
		)
	}
}
