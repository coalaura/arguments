package arguments

import (
	"os"
	"strings"
)

var (
	arguments map[string]argument
	noName    []argument
)

// I don't like golang flags package
func init() {
	arguments = make(map[string]argument)
	noName = make([]argument, 0)

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

					arguments[arg[2:2+index]] = argument{
						value: val,
					}
				} else {
					arguments[arg[2:]] = argument{}
				}

				name = ""
			} else {
				name = arg[1:]
			}
		} else {
			if name == "" {
				if arg != "" {
					noName = append(noName, argument{
						value: arg,
					})
				}
			} else {
				arguments[name] = argument{
					value: arg,
				}
			}

			name = ""
		}
	}

	if name != "" {
		arguments[name] = argument{}
	}
}

func getNamed(short, long string) argument {
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

func getUnnamed(index int) argument {
	if index < 0 || index >= len(noName) {
		return argument{
			isNil: true,
		}
	}

	return noName[index]
}

// IsNamedSet checks if the argument with the given short or long name is set.
// An argument is considered set if it is present in the command line arguments,
// even if it doesn't have a value.
func IsNamedSet(short, long string) bool {
	return !getNamed(short, long).isNil
}

// IsUnnamedSet checks if the unnamed argument with the given index is set.
// An unnamed argument is considered set if it is present in the command line arguments,
// even if it doesn't have a value.
func IsUnnamedSet(index int) bool {
	return !getUnnamed(index).isNil
}

// Get returns the raw value of the argument with the given short or long name.
// If the argument is not present, an empty string is returned.
func GetNamed(short, long string) string {
	return getNamed(short, long).value
}

// GetUnnamed returns the raw value of the unnamed argument with the given index.
// If the argument is not present, an empty string is returned.
func GetUnnamed(index int) string {
	return getUnnamed(index).value
}

// GetNamedAs takes an arguments short and long name and a default value of type T, and returns the
// value of the argument as type T. If the argument is not present, the default value is returned.
// If the type is a boolean then only "false" or "0" are considered false.
// If options are provided, the value is checked to ensure it is within the range of the
// options. If the value is not within the range, it is clamped to the closest value that
// is within the range. (Only for integers and floats)
func GetNamedAs[T any](short, long string, def T, options ...Options[T]) T {
	return convert(getNamed(short, long), def, options...)
}

// GetUnnamedAs takes an index and a default value of type T, and returns the value of the
// unnamed argument at that index as type T. If the argument is not present, the default
// value is returned.
// If the type is a boolean then only "false" or "0" are considered false.
// If options are provided, the value is checked to ensure it is within the range of the
// options. If the value is not within the range, it is clamped to the closest value that
// is within the range. (Only for integers and floats)
func GetUnnamedAs[T any](index int, def T, options ...Options[T]) T {
	return convert(getUnnamed(index), def, options...)
}

// NamedFile opens a file with the name of the argument given by the short or long name.
// If the argument is not present, the default file is returned.
// The file is opened with the given flags and permissions.
// If the file cannot be opened, an error is returned.
func NamedFile(short, long string, flag int, perm os.FileMode, def *os.File) (*os.File, error) {
	return asFile(GetNamed(short, long), flag, perm, def)
}

// UnnamedFile opens a file at the given index.
// If the argument is not present, the default file is returned.
// The file is opened with the given flags and permissions.
// If the file cannot be opened, an error is returned.
func UnnamedFile(index int, flag int, perm os.FileMode, def *os.File) (*os.File, error) {
	return asFile(GetUnnamed(index), flag, perm, def)
}
