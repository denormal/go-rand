package rand

// Strng returns a random string of length runes, or an error if sufficient
// random bits are not available.
func String(length int) (string, error) {
	return StringIf(length, nil)
}

// StringIf returns a random string of length runes, comprising only those runes
// for which isrune returns true. An error is returned if sufficient
// random bits are not available.
func StringIf(length int, isrune func(rune) bool) (string, error) {
	_runes, _err := RunesIf(length, isrune)
	if _err != nil {
		return "", _err
	}
	defer ReturnRunes(_runes)

	return string(_runes), nil
}

// StringFrom returns a random string of length runes, comprising only those
// runes listed in the alphabet. An error is returned if sufficient random bits
// are not available.
func StringFrom(length int, alphabet string) (string, error) {
	_runes, _err := RunesFrom(length, []rune(alphabet))
	defer ReturnRunes(_runes)

	return string(_runes), _err
}
