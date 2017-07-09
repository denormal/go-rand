package rand

import (
	"math"
)

func Float32() (float32, error) {
	_u32, _err := Uint32()
	if _err != nil {
		return 0, _err
	}

	return math.Float32frombits(_u32), nil
}

func Float64() (float64, error) {
	_u64, _err := Uint64()
	if _err != nil {
		return 0, _err
	}

	return math.Float64frombits(_u64), nil
}
