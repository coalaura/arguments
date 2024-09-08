package arguments

import (
	"strconv"
)

type argument struct {
	isNil bool
	value string
}

type Options[T any] struct {
	Min T
	Max T
}

func convert[T any](a argument, def T, options ...Options[T]) T {
	var value T

	if a.isNil {
		return def
	}

	var option Options[T]

	if len(options) > 0 {
		option = options[0]
	}

	switch any(value).(type) {
	case string:
		return any(a.value).(T)
	case []byte:
		return any([]byte(a.value)).(T)
	case bool:
		// If default is true, then only false and 0 are considered false
		if any(def).(bool) {
			return any(a.value != "false" && a.value != "0").(T)
		}

		// If default is false, then only true and 1 are considered true
		return any(a.value == "true" || a.value == "1").(T)
	case int64, int32, int16, int8, int:
		i, err := strconv.ParseInt(a.value, 10, 64)
		if err != nil {
			return def
		}

		return any(minmax(i, any(option).(Options[int64]))).(T)
	case uint64, uint32, uint16, uint8, uint, uintptr:
		i, err := strconv.ParseUint(a.value, 10, 64)
		if err != nil {
			return def
		}

		return any(minmax(i, any(option).(Options[uint64]))).(T)
	case float64, float32:
		i, err := strconv.ParseFloat(a.value, 64)
		if err != nil {
			return def
		}

		return any(minmax(i, any(option).(Options[float64]))).(T)
	}

	return value
}
