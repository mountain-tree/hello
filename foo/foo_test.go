package foo_test

import (
	"example/hello/foo"
	"testing"
)

func TestFoo(t *testing.T) {
	if foo.Bar() != "bar" {
		t.Fatal("Not a bar")
	}
}
