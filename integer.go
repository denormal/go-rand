package rand

func Uint() (uint, error) {
	_i, _err := Uint64()
	return uint(_i), _err
}

func Uint8() (uint8, error) {
	_byte, _err := Bytes(1)
	if _err != nil {
		return 0, _err
	}
	defer Put(_byte)

	return uint8(_byte[0]), nil
}

func Uint16() (uint16, error) {
	_bytes, _err := Bytes(2)
	if _err != nil {
		return 0, _err
	}
	defer Put(_bytes)

	return byteorder().Uint16(_bytes), nil
}

func Uint32() (uint32, error) {
	_bytes, _err := Bytes(4)
	if _err != nil {
		return 0, _err
	}
	defer Put(_bytes)

	return byteorder().Uint32(_bytes), nil
}

func Uint64() (uint64, error) {
	_bytes, _err := Bytes(8)
	if _err != nil {
		return 0, _err
	}
	defer Put(_bytes)

	return byteorder().Uint64(_bytes), nil
}

func Int() (int, error) {
	_i, _err := Int64()
	return int(_i), _err
}

func Int8() (int8, error) {
	_i, _err := Uint8()
	return int8(_i), _err
}

func Int16() (int16, error) {
	_i, _err := Uint16()
	return int16(_i), _err
}

func Int32() (int32, error) {
	_i, _err := Uint32()
	return int32(_i), _err
}

func Int64() (int64, error) {
	_i, _err := Uint64()
	return int64(_i), _err
}
