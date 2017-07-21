package rand

import (
	"math"
)

// Float32 returns a random 32 bit floating point value, or an error if
// sufficient random bits cannot be obtained by crypto/rand.
func Float32() (float32, error) {
	_u32, _err := UInt32()
	if _err != nil {
		return 0, _err
	}

	return math.Float32frombits(_u32), nil
}

// Float64 returns a random 64 bit floating point value, or an error if
// sufficient random bits cannot be obtained by crypto/rand.
func Float64() (float64, error) {
	_u64, _err := UInt64()
	if _err != nil {
		return 0, _err
	}

	return math.Float64frombits(_u64), nil
}
