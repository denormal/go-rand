package rand

// Uint() returns a random unsigned integer, or an error if sufficient random
// bits are not available to crypto/rand.
func Uint() (uint, error) {
	_i, _err := Uint64()
	return uint(_i), _err
}

// Uint8() returns a random unsigned 8 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Uint8() (uint8, error) {
	_byte, _err := Bytes(1)
	if _err != nil {
		return 0, _err
	}
	defer ReturnBytes(_byte)

	return uint8(_byte[0]), nil
}

// Uint16() returns a random unsigned 16 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Uint16() (uint16, error) {
	_bytes, _err := Bytes(2)
	if _err != nil {
		return 0, _err
	}
	defer ReturnBytes(_bytes)

	return byteorder().Uint16(_bytes), nil
}

// Uint32() returns a random unsigned 32 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Uint32() (uint32, error) {
	_bytes, _err := Bytes(4)
	if _err != nil {
		return 0, _err
	}
	defer ReturnBytes(_bytes)

	return byteorder().Uint32(_bytes), nil
}

// Uint64() returns a random unsigned 64 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Uint64() (uint64, error) {
	_bytes, _err := Bytes(8)
	if _err != nil {
		return 0, _err
	}
	defer ReturnBytes(_bytes)

	return byteorder().Uint64(_bytes), nil
}

// Int() returns a random signed integer, or an error if sufficient random
// bits are not available to crypto/rand.
func Int() (int, error) {
	_i, _err := Int64()
	return int(_i), _err
}

// Int8() returns a random signed 8 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Int8() (int8, error) {
	_i, _err := Uint8()
	return int8(_i), _err
}

// Int16() returns a random signed 16 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Int16() (int16, error) {
	_i, _err := Uint16()
	return int16(_i), _err
}

// Int32() returns a random signed 32 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Int32() (int32, error) {
	_i, _err := Uint32()
	return int32(_i), _err
}

// Int64() returns a random signed 64 bit integer, or an error if sufficient
// random bits are not available to crypto/rand.
func Int64() (int64, error) {
	_i, _err := Uint64()
	return int64(_i), _err
}

// Positive() returns a random, positive signed integer, or an error if
// sufficient random bits are not available to crypto/rand.
func Positive() (int, error) {
	_i, _err := Int()
	if _i < 0 {
		return -_i, _err
	} else {
		return _i, _err
	}
}

// Positive() returns a random, negative signed integer, or an error if
// sufficient random bits are not available to crypto/rand.
func Negative() (int, error) {
	_i, _err := Int()
	if _i < 0 {
		return _i, _err
	} else {
		return -_i, _err
	}
}
