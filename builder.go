package arguments

import "strings"

type Builder struct {
	strings.Builder

	colored bool
	reset   bool
}

// NewBuilder creates a new colored string builder
func NewBuilder(colored bool) *Builder {
	return &Builder{
		colored: colored,
		reset:   true,
	}
}

func (b *Builder) Mute() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;248m")

	b.reset = false
}

func (b *Builder) Text() {
	if !b.colored {
		return
	}

	b.WriteString("\033[3m\033[38;5;248m")

	b.reset = false
}

func (b *Builder) Name() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;153m")

	b.reset = false
}

func (b *Builder) Value() {
	if !b.colored {
		return
	}

	b.WriteString("\033[38;5;108m")

	b.reset = false
}

func (b *Builder) Reset() {
	if !b.colored || b.reset {
		return
	}

	b.WriteString("\033[0m")

	b.reset = true
}

func (b *Builder) String() string {
	b.Reset()

	return b.Builder.String()
}
