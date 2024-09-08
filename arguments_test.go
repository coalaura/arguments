package arguments

import (
	"os"
	"testing"
)

func TestArguments(t *testing.T) {
	os.Args = []string{
		"arguments.exe",
		"-u",
		"out.lar",
		"-v",
		"test",
	}

	Parse("u", "v")

	if !GetNamedAs("u", "unpack", false) {
		t.Error("unpack flag should be true")
	}

	if !GetNamedAs("v", "verbose", false) {
		t.Error("verbose flag should be true")
	}

	first := GetUnnamedAs(0, "")
	second := GetUnnamedAs(1, "")

	if first != "out.lar" {
		t.Errorf("first unnamed argument should be out.lar, got '%s'", first)
	}

	if second != "test" {
		t.Errorf("second unnamed argument should be test, got '%s'", second)
	}
}
