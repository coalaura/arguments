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
		"-a",
		"1,3,4,2",
		"one",
		"-n",
		"1234",
		"-bc",
		"0",
		"two",
		"-f",
		"123.56",
	}

	var (
		input    string
		output   string
		args     []uint16
		number   uint32
		boolean  bool
		boolean2 bool
		float    float64
	)

	Register("input", "i", &input)
	Register("output", "o", &output)
	Register("name", "n", &number)
	Register("args", "a", &args)
	Register("bool", "b", &boolean)
	Register("bool2", "c", &boolean2)
	Register("float", "f", &float)

	Parse()

	assertEqual(t, input, "input")
	assertEqual(t, output, "output")
	assertEqual(t, number, 1234)
	assertEqual(t, boolean, true)
	assertEqual(t, boolean2, false)
	assertEqual(t, float, 123.56)

	assertSliceEqual(t, args, []uint16{1, 3, 4, 2})
	assertSliceEqual(t, Args, []string{"one", "two"})
}

func assertEqual[V comparable](t *testing.T, actual, expected V) {
	if actual == expected {
		t.Logf("expected: %#v == %#v", expected, actual)
	} else {
		t.Errorf("expected: %#v != %#v", expected, actual)
	}
}

func assertSliceEqual[V comparable](t *testing.T, actual, expected []V) {
	if len(actual) != len(expected) {
		t.Errorf("expected: %#v != %#v", expected, actual)

		return
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected: %#v != %#v", expected, actual)

			return
		}
	}

	t.Logf("expected: %#v == %#v", expected, actual)
}
