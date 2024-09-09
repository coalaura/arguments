package arguments

import (
	"os"
	"strconv"
	"strings"
)

var (
	arguments map[string]argument
)

func init() {
	Parse()
}

func Parse() {
	arguments = make(map[string]argument)

	var (
		arg   string
		name  string
		val   string
		index int
	)

	for i := 1; i < len(os.Args); i++ {
		arg = os.Args[i]

		if arg[0] == '-' && len(arg) > 1 {
			if arg[1] == '-' {
				index = strings.Index(arg[2:], "=")

				if index >= 0 {
					val = ""

					if index+1 < len(arg) {
						val = arg[2+index+1:]
					}

					name = arg[2 : 2+index]

					if name != "" {
						arguments[name] = argument{
							value: val,
						}

						name = ""
					}
				} else {
					arguments[arg[2:]] = argument{}
				}

				name = ""
			} else {
				name = arg[1:]
			}
		} else if name != "" {
			arguments[name] = argument{
				value: arg,
			}

			name = ""
		}
	}

	if name != "" {
		arguments[name] = argument{}
	}
}

func get(short, long string) argument {
	arg, ok := arguments[short]

	if !ok && long != short {
		arg, ok = arguments[long]
	}

	if !ok {
		return argument{
			isNil: true,
		}
	}

	return arg
}

// IsNamedSet checks if the argument with the given short or long name is set.
// An argument is considered set if it is present in the command line arguments,
// even if it doesn't have a value.
func IsSet(short, long string) bool {
	return !get(short, long).isNil
}

// String returns the value of the string argument with the given short or long name.
// If the argument is not present, the default value is returned.
func String(short, long, def string) string {
	arg := get(short, long)

	if arg.isNil {
		return def
	}

	return arg.value
}

// Bool returns the value of the boolean argument with the given short or long name.
// If the argument is not present, the default value is returned.
// The function considers the argument to be true if it is present and its value is
// not "false" or "0".
func Bool(short, long string, def bool) bool {
	arg := get(short, long)

	if arg.isNil {
		return def
	}

	return arg.value != "false" && arg.value != "0"
}

// IntN returns the value of the integer argument with the given short or long name.
// If the argument is not present, the default value is returned.
// The function supports the following types: int64, int32, int16, int8, int.
// The function will return the default value if the argument is not a valid integer.
func IntN[T int64 | int32 | int16 | int8 | int](short, long string, def T, options ...Options[T]) T {
	val := get(short, long).value

	if val == "" {
		return def
	}

	var bits int

	switch any(def).(type) {
	case int64, int:
		bits = 64
	case int32:
		bits = 32
	case int16:
		bits = 16
	case int8:
		bits = 8
	}

	i, err := strconv.ParseInt(val, 10, bits)
	if err != nil {
		return def
	}

	v := T(i)

	if len(options) > 0 {
		min := options[0].Min
		max := options[0].Max

		if v < min {
			return min
		}

		if v > max {
			return max
		}
	}

	return v
}

// UIntN returns the value of the unsigned integer argument with the given short or long name.
// If the argument is not present, the default value is returned.
// The function supports the following types: uint64, uint32, uint16, uint8, uint, uintptr.
// The function will return the default value if the argument is not a valid unsigned integer.
func UIntN[T uint64 | uint32 | uint16 | uint8 | uint | uintptr](short, long string, def T, options ...Options[T]) T {
	val := get(short, long).value

	if val == "" {
		return def
	}

	var bits int

	switch any(def).(type) {
	case uint64, uint, uintptr:
		bits = 64
	case uint32:
		bits = 32
	case uint16:
		bits = 16
	case uint8:
		bits = 8
	}

	u, err := strconv.ParseUint(val, 10, bits)
	if err != nil {
		return def
	}

	v := T(i)

	if len(options) > 0 {
		min := options[0].Min
		max := options[0].Max

		if v < min {
			return min
		}

		if v > max {
			return max
		}
	}

	return v
}

// FloatN returns the value of the float argument with the given short or long name.
// If the argument is not present, the default value is returned.
// The function supports the following types: float64, float32.
// The function will return the default value if the argument is not a valid float.
func FloatN[T float64 | float32](short, long string, def T, options ...Options[T]) T {
	val := get(short, long).value

	if val == "" {
		return def
	}

	var bits int

	switch any(def).(type) {
	case float64:
		bits = 64
	case float32:
		bits = 32
	}

	f, err := strconv.ParseFloat(val, bits)
	if err != nil {
		return def
	}

	v := T(i)

	if len(options) > 0 {
		min := options[0].Min
		max := options[0].Max

		if v < min {
			return min
		}

		if v > max {
			return max
		}
	}

	return v
}

// NamedFile opens a file with the name of the argument given by the short or long name.
// If the argument is not present, the default file is returned.
// The file is opened with the given flags and permissions.
// If the file cannot be opened, an error is returned.
func File(short, long string, flag int, perm os.FileMode, def *os.File) (*os.File, error) {
	return asFile(get(short, long).value, flag, perm, def)
}
