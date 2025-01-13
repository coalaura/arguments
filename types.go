package arguments

import (
	"unsafe"
)

type value interface {
	string | bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | []string | []int | []int8 | []int16 | []int32 | []int64 | []uint | []uint8 | []uint16 | []uint32 | []uint64 | []uintptr | []float32 | []float64
}

type setter interface {
	Set(string)
}

type holder[V any] struct {
	long  string
	short string
	value *V
}

type container struct {
	short map[string]setter
	long  map[string]setter
}

func (h *holder[V]) Set(value string) {
	if value == "" {
		// Special handling for bool type
		if _, ok := any(*h.value).(bool); ok {
			*(*bool)(unsafe.Pointer(h.value)) = true
		}

		return
	}

	*h.value = as(value, *h.value)
}

func (c *container) SetShort(short string, value string) {
	arg, ok := c.short[short]
	if !ok {
		return
	}

	arg.Set(value)
}

func (c *container) SetLong(long string, value string) {
	arg, ok := c.long[long]
	if !ok {
		return
	}

	arg.Set(value)
}
