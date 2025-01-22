package arguments

import (
	"fmt"
	"os"
	"strings"
)

const (
	length = 80
	indent = "     " // 4 + 1 space
)

// WithHelp sets the help for the argument
func (h *holder[V]) WithHelp(text string) {
	if len(text) <= length {
		h.text = text

		return
	}

	fields := strings.Fields(text)

	var (
		b    strings.Builder
		line int
	)

	for _, field := range fields {
		if line+len(field)+1 > length {
			b.WriteRune('\n')
			b.WriteString(indent)

			line = 0
		} else if line > 0 {
			b.WriteRune(' ')
			line++
		}

		b.WriteString(field)

		line += len(field)
	}

	h.text = b.String()
}

func (h *holder[V]) write(b *Builder) {
	typeName := fmt.Sprintf("%T", *h.value)
	isBoolean := typeName == "bool"

	if h.short != "\x00" {
		b.Mute()
		b.WriteString(" -")

		b.Name()
		b.WriteString(h.short)

		if !isBoolean {
			b.WriteRune(' ')

			b.Value()
			b.WriteString(typeName)
		}
	}

	if h.long != "" {
		b.Mute()

		if h.short != "\x00" {
			b.WriteRune(',')
		}

		b.WriteString(" --")

		b.Name()
		b.WriteString(h.long)

		if !isBoolean {
			b.Mute()
			b.WriteRune('=')

			b.Value()
			b.WriteString(typeName)
		}
	}

	if h.text != "" {
		b.Text()

		b.WriteRune('\n')
		b.WriteString(indent)

		b.WriteString(h.text)
	}

	b.Reset()
}

// ShowHelp displays the help for all arguments
func ShowHelp(colored bool) {
	help := NewBuilder(colored)

	for _, arg := range arguments.list {
		if help.Len() > 0 {
			help.WriteString("\n")
		}

		arg.write(help)
	}

	fmt.Println(help.String())
}

// ShowHelpAndExit displays the help for all arguments and then exits
func ShowHelpAndExit(colored bool) {
	ShowHelp(colored)

	os.Exit(0)
}
