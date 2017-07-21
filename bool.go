package rand

// Bool returns a random boolean value, or an error if the underlying call
// to crypto/rand fails.
func Bool() (bool, error) {
	_i, _err := Int8()
	return _i < 0, _err
}
