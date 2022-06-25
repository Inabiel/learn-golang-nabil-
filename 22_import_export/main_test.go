package main

import (
	"impex/folder1"
	"testing"
)

func Test(t *testing.T) {
	want := "Hello!"
	if got := folder1.SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
