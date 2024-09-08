package arguments

import "golang.org/x/exp/constraints"

func minmax[T constraints.Integer | constraints.Float](val T, options ...Options[T]) T {
	// No options means we don't need to do anything
	if len(options) == 0 {
		return val
	}

	min := options[0].Min
	max := options[0].Max

	if min != 0 && val < min {
		return min
	}

	if max != 0 && val > max {
		return max
	}

	return val
}
