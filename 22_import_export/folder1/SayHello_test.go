package folder1

import (
	"testing"
)

func Test(t *testing.T) {
	want := "Hello!"
	if got := SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
