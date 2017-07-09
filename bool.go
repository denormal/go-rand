package rand

func Bool() (bool, error) {
	_u, _err := Uint8()
	return (_u % 2) == 0, _err
}
