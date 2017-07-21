package rand_test

import (
	"testing"
	"time"

	"github.com/denormal/go-rand"
)

func TestTime(t *testing.T) {
	var (
		_zero int
		_nil  time.Time
	)

	for _i := 0; _i < MAXCOUNT; _i++ {
		_time, _err := rand.Time()
		if _err != nil {
			t.Fatalf("Unexpected error from rand.Time(): %s", _err)
		} else if _nil.Equal(_time) {
			_zero++
		}
	}

	if _zero == MAXCOUNT {
		t.Fatal("Expected non-zero times from rand.Time(); none returned")
	}
}

func TestDuration(t *testing.T) {
	var (
		_zero int
		_nil  time.Duration
	)

	for _i := 0; _i < MAXCOUNT; _i++ {
		_duration, _err := rand.Duration()
		if _err != nil {
			t.Fatalf("Unexpected error from rand.Duration(): %s", _err)
		} else if _duration == _nil {
			_zero++
		}
	}

	if _zero == MAXCOUNT {
		t.Fatal("Expected non-zero times from rand.Duration(); none returned")
	}
}
