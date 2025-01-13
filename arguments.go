package arguments

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
)

var (
	parsed atomic.Int32

	arguments = container{
		shorts: make(map[rune]setter),
		longs:  make(map[string]setter),
	}

	// Args contains the remaining unnamed arguments
	Args []string
)

// RegisterHelp registers the help argument (-h or --help) which will display the help for all arguments after parsing and then exit
func RegisterHelp(colored bool, text ...string) {
	if colored {
		arguments.colored = true
	}

	var value bool
	arguments.help = Register("help", 'h', &value)

	if len(text) > 0 {
		arguments.help.WithHelp(text[0])
	}
}

// Register registers an argument with the given short and long names and value pointer
func Register[V value](long string, short rune, value *V) *holder[V] {
	if short == 0 && long == "" {
		panic("either short or long name must be set")
	}

	arg := &holder[V]{
		long:  long,
		short: string(short),
		value: value,
	}

	arguments.list = append(arguments.list, arg)

	if short != 0 {
		if _, ok := arguments.shorts[short]; ok {
			panic(fmt.Sprintf("argument '-%c' already registered", short))
		}

		arguments.shorts[short] = arg
	}

	if long != "" {
		if _, ok := arguments.longs[long]; ok {
			panic(fmt.Sprintf("argument '--%s' already registered", long))
		}

		arguments.longs[long] = arg
	}

	return arg
}

// Parse parses the command line arguments and sets the values of the arguments (can only be called once)
func Parse() error {
	if !parsed.CompareAndSwap(0, 1) {
		return errors.New("arguments already parsed")
	}

	var (
		index int
		name  rune
	)

	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			if name != 0 {
				arguments.short(name, "")

				name = 0
			}

			// --argument=value
			if len(arg) > 1 && arg[1] == '-' {
				index = strings.Index(arg, "=")

				if index == -1 {
					arguments.long(arg[2:], "")
				} else {
					arguments.long(arg[2:index], arg[index+1:])
				}

				continue
			}

			// -a value
			for x, rn := range arg[1:] {
				if rn == ' ' || rn == '=' {
					continue
				}

				// -abcdefg
				if x > 0 && name != 0 {
					arguments.short(name, "")
				}

				name = rn
			}
		} else if name != 0 {
			arguments.short(name, arg)

			name = 0
		} else {
			Args = append(Args, arg)
		}
	}

	if name != 0 {
		arguments.short(name, "")
	}

	if arguments.help != nil && *arguments.help.value {
		ShowHelpAndExit(arguments.colored)
	}

	if arguments.invalid != "" {
		return fmt.Errorf("invalid argument: %s", arguments.invalid)
	}

	return nil
}

// MustParse parses the command line arguments and sets the values of the arguments, it will print an error message and exit if an error occurs
func MustParse() {
	if err := Parse(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}
}
