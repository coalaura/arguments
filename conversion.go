package arguments

import (
	"strconv"
	"strings"
)

func as[V any](raw string, def V) V {
	switch any(def).(type) {
	case string:
		return any(raw).(V)
	case bool:
		return any(raw != "false" && raw != "0").(V)

	// Integers
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		i, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return def
		}

		switch any(def).(type) {
		case int:
			return any(int(i)).(V)
		case int8:
			return any(int8(i)).(V)
		case int16:
			return any(int16(i)).(V)
		case int32:
			return any(int32(i)).(V)
		case uint:
			return any(uint(i)).(V)
		case uint8:
			return any(uint8(i)).(V)
		case uint16:
			return any(uint16(i)).(V)
		case uint32:
			return any(uint32(i)).(V)
		case uint64:
			return any(uint64(i)).(V)
		case uintptr:
			return any(uintptr(i)).(V)
		}

		return any(i).(V)

	// Floats
	case float32, float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return def
		}

		switch any(def).(type) {
		case float32:
			return any(float32(f)).(V)
		}

		return any(f).(V)

	// Slices
	case []string, []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []uintptr, []float32, []float64:
		s := strings.Split(raw, ",")

		switch any(def).(type) {
		case []int:
			r := make([]int, len(s))

			for i, v := range s {
				r[i] = any(as(v, 0)).(int)
			}

			return any(r).(V)
		case []int8:
			r := make([]int8, len(s))

			for i, v := range s {
				r[i] = any(as[int8](v, 0)).(int8)
			}

			return any(r).(V)
		case []int16:
			r := make([]int16, len(s))

			for i, v := range s {
				r[i] = any(as[int16](v, 0)).(int16)
			}

			return any(r).(V)
		case []int32:
			r := make([]int32, len(s))

			for i, v := range s {
				r[i] = any(as[int32](v, 0)).(int32)
			}

			return any(r).(V)
		case []int64:
			r := make([]int64, len(s))

			for i, v := range s {
				r[i] = any(as[int64](v, 0)).(int64)
			}

			return any(r).(V)
		case []uint:
			r := make([]uint, len(s))

			for i, v := range s {
				r[i] = any(as[uint](v, 0)).(uint)
			}

			return any(r).(V)
		case []uint8:
			r := make([]uint8, len(s))

			for i, v := range s {
				r[i] = any(as[uint8](v, 0)).(uint8)
			}

			return any(r).(V)
		case []uint16:
			r := make([]uint16, len(s))

			for i, v := range s {
				r[i] = any(as[uint16](v, 0)).(uint16)
			}

			return any(r).(V)
		case []uint32:
			r := make([]uint32, len(s))

			for i, v := range s {
				r[i] = any(as[uint32](v, 0)).(uint32)
			}

			return any(r).(V)
		case []uint64:
			r := make([]uint64, len(s))

			for i, v := range s {
				r[i] = any(as[uint64](v, 0)).(uint64)
			}

			return any(r).(V)
		case []uintptr:
			r := make([]uintptr, len(s))

			for i, v := range s {
				r[i] = any(as[uintptr](v, 0)).(uintptr)
			}

			return any(r).(V)
		case []float32:
			r := make([]float32, len(s))

			for i, v := range s {
				r[i] = any(as[float32](v, 0)).(float32)
			}

			return any(r).(V)
		case []float64:
			r := make([]float64, len(s))

			for i, v := range s {
				r[i] = any(as[float64](v, 0)).(float64)
			}

			return any(r).(V)
		}

		return any(s).(V)
	}

	panic("unsupported type")
}
