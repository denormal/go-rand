package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestComplex64(t *testing.T) {
	numeric(
		t,
		"Complex64",
		func() (interface{}, error) { return rand.Complex64() },
	)
}

func TestComplex128(t *testing.T) {
	numeric(
		t,
		"Complex128",
		func() (interface{}, error) { return rand.Complex128() },
	)
}
