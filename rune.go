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

func Runes(length int) ([]rune, error) {
	return RunesIf(length, nil)
}

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
			Return(_runes)
			return nil, _err
		}
		defer Put(_rand)

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

func RunesFrom(length int, runes []rune) ([]rune, error) {
	if length == 0 {
		return []rune{}, nil
	}

	_runes := rpool.Get().([]rune)
	if cap(_runes) < length {
		_runes = make([]rune, 0, length)
	}

	_alphabet := len(runes)
	for _i := 0; _i < length; _i++ {
		_index, _err := Int()
		if _err != nil {
			Return(_runes)
			return nil, _err
		}

		_runes[_i] = runes[_index%_alphabet]
	}

	return _runes, nil
}

func Return(runes []rune) {
	if cap(runes) != 0 {
		rpool.Put(runes[:0])
	}
}
