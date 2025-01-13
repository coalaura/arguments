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

	Register("input", 'i', &input).WithHelp("The input file")
	Register("output", 0, &output).WithHelp("The output file")
	Register("number", 'n', &number).WithHelp("Number of things")
	Register("args", 'a', &args).WithHelp("Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus.")
	Register("bool", 'b', &boolean).WithHelp("Boolean thats true")
	Register("bool2", 'c', &boolean2).WithHelp("Boolean thats false")
	Register("float", 'f', &float).WithHelp("Floaty thingy")

	Parse()

	ShowHelp(true)

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
