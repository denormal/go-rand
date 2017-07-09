package rand

import (
	"time"
)

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

func Duration() (time.Duration, error) {
	_i, _err := Int64()
	if _err == nil {
		return time.Duration(_i), nil
	} else {
		return time.Duration(0), _err
	}
}
