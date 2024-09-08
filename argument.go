package arguments

import "strconv"

type Argument struct {
	name string

	IsNil bool
	Value string
}

type Options[V int | int64 | uint64 | float64] struct {
	Min V
	Max V
}

// IsSet returns true if the argument is set, false otherwise.
// It is set if it exists in the command line arguments even if it has no value.
func (a Argument) IsSet() bool {
	return !a.IsNil
}

// String returns the value of the argument if set, otherwise the default value.
func (a Argument) String(def string) string {
	if a.IsNil {
		return def
	}

	return a.Value
}

// Bool returns the value of the argument as a boolean if set, otherwise the default value.
// The value is considered true if it is not "false" or "0".
func (a Argument) Bool(def bool) bool {
	if a.IsNil {
		return def
	}

	if a.Value == "false" || a.Value == "0" {
		return false
	}

	return true
}

// Int returns the value of the argument as an int if set, otherwise the default value.
// If the value is not an integer (can't be parsed as an int), the default value is returned.
// If any options are given, the value is clamped to the options' range.
func (a Argument) Int(def int, options ...Options[int]) int {
	if a.IsNil {
		return def
	}

	i, err := strconv.Atoi(a.Value)
	if err != nil {
		return def
	}

	return minmax(i, options...)
}

// Int64 returns the value of the argument as an int64 if set, otherwise the default value.
// If the value is not an integer (can't be parsed as an int64), the default value is returned.
// If any options are given, the value is clamped to the options' range.
func (a Argument) Int64(def int64, options ...Options[int64]) int64 {
	if a.IsNil {
		return def
	}

	i, err := strconv.ParseInt(a.Value, 10, 64)
	if err != nil {
		return def
	}

	return minmax(i, options...)
}

// Uint64 returns the value of the argument as a uint64 if set, otherwise the default value.
// If the value is not an unsigned integer (can't be parsed as a uint64), the default value is returned.
// If any options are given, the value is clamped to the options' range.
func (a Argument) Uint64(def uint64, options ...Options[uint64]) uint64 {
	if a.IsNil {
		return def
	}

	i, err := strconv.ParseUint(a.Value, 10, 64)
	if err != nil {
		return def
	}

	return minmax(i, options...)
}

// Float64 returns the value of the argument as a float64 if set, otherwise the default value.
// If the value is not a floating point number (can't be parsed as a float64), the default value is returned.
// If any options are given, the value is clamped to the options' range.
func (a Argument) Float64(def float64, options ...Options[float64]) float64 {
	if a.IsNil {
		return def
	}

	i, err := strconv.ParseFloat(a.Value, 64)
	if err != nil {
		return def
	}

	return minmax(i, options...)
}
