package arguments

import (
	"os"
	"strings"
)

type Arguments struct {
	Arguments map[string]Argument
}

var (
	arguments Arguments
)

// I don't like golang flags package
func init() {
	arguments = Arguments{
		Arguments: make(map[string]Argument),
	}

	var (
		arg   string
		val   string
		index int

		current Argument
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

					arguments.set(Argument{
						name:  arg[2 : 2+index],
						Value: val,
					})
				} else {
					arguments.set(Argument{
						name: arg[2:],
					})
				}

				current = Argument{}
			} else {
				current = Argument{
					name: arg[1:],
				}
			}
		} else {
			current.Value = arg

			arguments.set(current)

			current = Argument{}
		}
	}

	if current.name != "" {
		arguments.set(current)
	}
}

func (a *Arguments) set(arg Argument) {
	a.Arguments[arg.name] = arg
}

func (a *Arguments) Arg(short, long string) Argument {
	arg, ok := a.Arguments[short]

	if !ok && long != short {
		arg, ok = a.Arguments[long]
	}

	if !ok {
		return Argument{
			IsNil: true,
		}
	}

	return arg
}

func (a *Arguments) Has(short, long string) bool {
	return a.Arg(short, long).IsSet()
}

func (a *Arguments) String(short, long, def string) string {
	return a.Arg(short, long).String(def)
}

func (a *Arguments) Bool(short, long string, def bool) bool {
	return a.Arg(short, long).Bool(def)
}

func (a *Arguments) Int(short, long string, def int, options ...Options[int]) int {
	return a.Arg(short, long).Int(def, options...)
}

func (a *Arguments) Int64(short, long string, def int64, options ...Options[int64]) int64 {
	return a.Arg(short, long).Int64(def, options...)
}

func (a *Arguments) Uint64(short, long string, def uint64, options ...Options[uint64]) uint64 {
	return a.Arg(short, long).Uint64(def, options...)
}

func (a *Arguments) Float64(short, long string, def float64, options ...Options[float64]) float64 {
	return a.Arg(short, long).Float64(def, options...)
}
