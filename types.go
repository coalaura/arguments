package arguments

import (
	"fmt"
	"unsafe"
)

type value interface {
	string | bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | []string | []int | []int8 | []int16 | []int32 | []int64 | []uint | []uint8 | []uint16 | []uint32 | []uint64 | []uintptr | []float32 | []float64
}

type setter interface {
	set(string)
	write(*builder)
}

type holder[V any] struct {
	text string

	long  string
	short string
	value *V
}

type container struct {
	colored bool
	help    *holder[bool]

	invalid string

	list   []setter
	shorts map[rune]setter
	longs  map[string]setter
}

func (h *holder[V]) set(value string) {
	if value == "" {
		// Special handling for bool type
		if _, ok := any(*h.value).(bool); ok {
			*(*bool)(unsafe.Pointer(h.value)) = true
		}

		return
	}

	*h.value = as(value, *h.value)
}

func (c *container) short(short rune, value string) {
	arg, ok := c.shorts[short]
	if !ok {
		c.invalid = fmt.Sprintf("-%c", short)

		return
	}

	arg.set(value)
}

func (c *container) long(long string, value string) {
	arg, ok := c.longs[long]
	if !ok {
		c.invalid = fmt.Sprintf("--%s", long)

		return
	}

	arg.set(value)
}
