package rand

import (
	"time"
)

// Time returns a random time, or an error if it is not possible to read
// sufficient random bits.
func Time() (time.Time, error) {
	var (
		_err    error
		_s, _us int64
	)

	_s, _err = Int64()
	if _err == nil {
		_us, _err = Int64()
		if _err == nil {
			return time.Unix(_s, _us), nil
		}
	}

	return time.Time{}, _err
}

// Time returns a random duration, or an error if it is not possible to read
// sufficient random bits.
func Duration() (time.Duration, error) {
	_i, _err := Int64()
	return time.Duration(_i), _err
}
