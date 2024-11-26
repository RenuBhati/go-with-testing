package integers

import (
	"fmt"
	"testing"
)

func Adder(a, b int) int {
	return a + b
}
func ExampleAdd() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
func TestExampleAdd(t *testing.T) {
	sum := Adder(1, 5)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
