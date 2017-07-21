package rand

import (
	"sync"
	"unicode/utf8"
)

var (
	size  = 256
	rpool = sync.Pool{
		New: func() interface{} { return make([]rune, 0, length) },
	}
)

// Runes returns a random slice of length runes, or an error if sufficient
// random bits are not available. The rune slice is taken from a pool of slices.
func Runes(length int) ([]rune, error) {
	return RunesIf(length, nil)
}

// RunesIf returns a random slice of length runes, comprising only those runes
// for which isrune returns true. An error is returned if sufficient
// random bits are not available. The rune slice is taken from a pool of slices.
func RunesIf(length int, isrune func(rune) bool) ([]rune, error) {
	if length <= 0 {
		return []rune{}, nil
	}

	_runes := rpool.Get().([]rune)
	if cap(_runes) < length {
		_runes = make([]rune, 0, length)
	}

	for {
		_rand, _err := Bytes(length)
		if _err != nil {
			ReturnRunes(_runes)
			return nil, _err
		}
		defer ReturnBytes(_rand)

		_bytes := _rand
		for len(_bytes) != 0 {
			_rune, _length := utf8.DecodeRune(_bytes)
			if _rune != utf8.RuneError {
				if isrune == nil || isrune(_rune) {
					_runes = append(_runes, _rune)
					if len(_runes) == length {
						return _runes, nil
					}
				}
			}

			_bytes = _bytes[_length:]
		}
	}
}

// RunesFrom returns a random slice of length runes, comprising only those runes
// listed in the alphabet slice. An error is returned if sufficient random bits
// are not available. The rune slice is taken from a pool of slices.
func RunesFrom(length int, alphabet []rune) ([]rune, error) {
	if length <= 0 {
		return []rune{}, nil
	}

	_runes := rpool.Get().([]rune)
	if cap(_runes) < length {
		_runes = make([]rune, 0, length)
	}
	_runes = _runes[:length]

	_alphabet := len(alphabet)
	for _i := 0; _i < length; _i++ {
		_index, _err := Positive()
		if _err != nil {
			ReturnRunes(_runes)
			return nil, _err
		}

		_runes[_i] = alphabet[_index%_alphabet]
	}

	return _runes, nil
}

// ReturnRunes returns a rune slice to the pool of slices that may be returned
// by Runes(), RunesIf() or RunesFrom().
func ReturnRunes(runes []rune) {
	if cap(runes) != 0 {
		rpool.Put(runes[:0])
	}
}
