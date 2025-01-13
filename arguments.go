package arguments

import (
	"os"
	"strings"
	"sync/atomic"
)

var (
	parsed atomic.Int32

	arguments = container{
		short: make(map[string]setter),
		long:  make(map[string]setter),
	}

	// Args contains the remaining unnamed arguments
	Args []string
)

func Register[V value](long, short string, value *V, def ...V) {
	arg := &holder[V]{
		long:  long,
		short: short,
		value: value,
	}

	if len(def) > 0 {
		*value = def[0]
	}

	arguments.short[short] = arg
	arguments.long[long] = arg
}

func Parse() {
	if !parsed.CompareAndSwap(0, 1) {
		return
	}

	var (
		index int
		name  rune
	)

	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			name = 0

			// --argument=value
			if len(arg) > 1 && arg[1] == '-' {
				index = strings.Index(arg, "=")

				if index == -1 {
					arguments.SetLong(arg[2:], "")
				} else {
					arguments.SetLong(arg[2:index], arg[index+1:])
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
					arguments.SetShort(string(name), "")
				}

				name = rn
			}
		} else if name != 0 {
			arguments.SetShort(string(name), arg)

			name = 0
		} else {
			Args = append(Args, arg)
		}
	}
}
