package testhelpers

import (
	"fmt"
	"testing"
)

func Given(s string) {
	fmt.Println("Given", s)
}

func When(s string) {
	fmt.Println("  When", s)
}

func Then(s string) {
	fmt.Println("    Then", s)
}

func And(s string) {
	fmt.Println("    And", s)
}

func Assert(a, x interface{}, t *testing.T) {

	a = fmt.Sprintf("%v", a)
	x = fmt.Sprintf("%v", x)
	if a != x {
		t.Errorf("Expected %s, but got: %s", x, a)
	}
}
