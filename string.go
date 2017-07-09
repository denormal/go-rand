package rand

func String(length int) (string, error) {
	return StringIf(length, nil)
}

func StringIf(length int, cmp func(rune) bool) (string, error) {
	_runes, _err := RunesIf(length, cmp)
	if _err != nil {
		return "", _err
	}
	defer Return(_runes)

	return string(_runes), nil
}

func StringFrom(length int, runes []rune) (string, error) {
	_runes, _err := RunesFrom(length, runes)
	defer Return(_runes)

	return string(_runes), _err
}
