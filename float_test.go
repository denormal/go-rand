package rand_test

import (
	"testing"

	"github.com/denormal/go-rand"
)

func TestFloat32(t *testing.T) {
	numeric(t, "Float32", func() (interface{}, error) { return rand.Float32() })
}

func TestFloat64(t *testing.T) {
	numeric(t, "Float64", func() (interface{}, error) { return rand.Float64() })
}
