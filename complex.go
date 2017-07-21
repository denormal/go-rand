package rand

// Complex64 returns a random 64 bit complex value, or an error if insufficient
// random bits can be obtained.
func Complex64() (complex64, error) {
	var (
		_r   float32
		_i   float32
		_err error
		_nil complex64
	)

	_r, _err = Float32()
	if _err == nil {
		_i, _err = Float32()
		if _err == nil {
			return complex(_r, _i), nil
		}
	}

	return _nil, _err
}

// Complex128 returns a random 128 bit complex value, or an error if
// insufficient random bits can be obtained.
func Complex128() (complex128, error) {
	var (
		_r   float64
		_i   float64
		_err error
		_nil complex128
	)

	_r, _err = Float64()
	if _err == nil {
		_i, _err = Float64()
		if _err == nil {
			return complex(_r, _i), nil
		}
	}

	return _nil, _err
}
