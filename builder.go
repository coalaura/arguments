package arguments

import "strings"

type builder struct {
	colored bool
	strings.Builder
}

func (b *builder) mute() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;248m")
}

func (b *builder) text() {
	if !b.colored {
		return
	}

	b.WriteString("\033[3m\033[38;5;248m")
}

func (b *builder) name() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;153m")
}

func (b *builder) value() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;108m")
}

func (b *builder) reset() {
	if !b.colored {
		return
	}

	b.WriteString("\033[0m")
}
