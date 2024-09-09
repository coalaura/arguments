package arguments

import (
	"os"
	"testing"
)

func TestArguments(t *testing.T) {
	os.Args = []string{
		"arguments.exe",
		"-i",
		"input",
		"--output=output",
		"-n",
		"1234",
		"-f",
		"123.56",
		"-b",
	}

	Parse()

	i := String("i", "input", "")
	o := String("o", "output", "")
	n := IntN("n", "number", 0)
	f := FloatN("f", "float", 0.0)
	b := Bool("b", "bool", false)

	if i != "input" {
		t.Errorf("expected 'input', got '%s'", i)
	}

	if o != "output" {
		t.Errorf("expected 'output', got '%s'", o)
	}

	if n != 1234 {
		t.Errorf("expected 1234, got '%d'", n)
	}

	if f != 123.56 {
		t.Errorf("expected 123.56, got '%f'", f)
	}

	if !b {
		t.Errorf("expected true, got '%t'", b)
	}
}
