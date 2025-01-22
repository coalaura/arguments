package arguments

import "strings"

type Builder struct {
	strings.Builder
	colored bool
}

// NewBuilder creates a new colored string builder
func NewBuilder(colored bool) *Builder {
	return &Builder{
		colored: colored,
	}
}

func (b *Builder) mute() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;248m")
}

func (b *Builder) text() {
	if !b.colored {
		return
	}

	b.WriteString("\033[3m\033[38;5;248m")
}

func (b *Builder) name() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;153m")
}

func (b *Builder) value() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;108m")
}

func (b *Builder) reset() {
	if !b.colored {
		return
	}

	b.WriteString("\033[0m")
}
