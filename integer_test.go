package rand_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/denormal/go-rand"
)

func TestInt(t *testing.T) {
	numeric(t, "Int", func() (interface{}, error) { return rand.Int() })
}

func TestInt8(t *testing.T) {
	numeric(t, "Int8", func() (interface{}, error) { return rand.Int8() })
}

func TestInt16(t *testing.T) {
	numeric(t, "Int16", func() (interface{}, error) { return rand.Int16() })
}

func TestInt32(t *testing.T) {
	numeric(t, "Int32", func() (interface{}, error) { return rand.Int32() })
}

func TestInt64(t *testing.T) {
	numeric(t, "Int64", func() (interface{}, error) { return rand.Int64() })
}

func TestUint(t *testing.T) {
	numeric(t, "Uint", func() (interface{}, error) { return rand.Uint() })
}

func TestUint8(t *testing.T) {
	numeric(t, "Uint8", func() (interface{}, error) { return rand.Uint8() })
}

func TestUint16(t *testing.T) {
	numeric(t, "Uint16", func() (interface{}, error) { return rand.Uint16() })
}

func TestUint32(t *testing.T) {
	numeric(t, "Uint32", func() (interface{}, error) { return rand.Uint32() })
}

func TestUint64(t *testing.T) {
	numeric(t, "Uint64", func() (interface{}, error) { return rand.Uint64() })
}

func TestPositive(t *testing.T) {
	numericif(
		t,
		"Positive",
		func() (interface{}, error) { return rand.Positive() },
		func(v interface{}) error {
			switch _v := v.(type) {
			case int:
				if _v < 0 {
					return fmt.Errorf(
						"Expected positive value from rand.Positive(), got %d",
						_v,
					)
				}
			default:
				return fmt.Errorf(
					"Expected int from rand.Positive(), got %q",
					reflect.TypeOf(v),
				)
			}

			return nil
		},
	)
}

func TestNegative(t *testing.T) {
	numericif(
		t,
		"Negative",
		func() (interface{}, error) { return rand.Negative() },
		func(v interface{}) error {
			switch _v := v.(type) {
			case int:
				if _v > 0 {
					return fmt.Errorf(
						"Expected negative value from rand.Negative(), got %d",
						_v,
					)
				}
			default:
				return fmt.Errorf(
					"Expected int from rand.Negative(), got %q",
					reflect.TypeOf(v),
				)
			}

			return nil
		},
	)
}
