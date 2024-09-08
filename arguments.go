package arguments

import (
	"os"
	"strings"
)

var (
	arguments map[string]Argument
)

// I don't like golang flags package
func init() {
	arguments = make(map[string]Argument)

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

					arguments[arg[2:2+index]] = Argument{
						value: val,
					}
				} else {
					arguments[arg[2:]] = Argument{}
				}

				name = ""
			} else {
				name = arg[1:]
			}
		} else {
			arguments[name] = Argument{
				value: arg,
			}

			name = ""
		}
	}

	if name != "" {
		arguments[name] = Argument{}
	}
}

func get(short, long string) Argument {
	arg, ok := arguments[short]

	if !ok && long != short {
		arg, ok = arguments[long]
	}

	if !ok {
		return Argument{
			isNil: true,
		}
	}

	return arg
}

// IsSet checks if the argument with the given short or long name is set.
// An argument is considered set if it is present in the command line arguments,
// even if it doesn't have a value.
func IsSet(short, long string) bool {
	return !get(short, long).isNil
}

// Get returns the raw value of the argument with the given short or long name.
// If the argument is not present, an empty string is returned.
func Get(short, long string) string {
	return get(short, long).value
}

// GetAs takes an Argument and a default value of type T, and returns the value of the Argument
// as type T. If the Argument is nil, the default value is returned. If the Argument is not nil,
// the value is attempted to be converted to type T. If the conversion fails, an error is returned.
// If the type is a boolean and the default value is true, then only "false" and "0" are considered
// false. If the default value is false, then only "true" and "1" are considered true.
// If options are provided, the value is checked to ensure it is within the range of the
// options. If the value is not within the range, it is clamped to the closest value that
// is within the range. (Only for integers and floats)
func GetAs[T any](short, long string, def T, options ...Options[T]) (T, error) {
	return convert(get(short, long), def, options...)
}
