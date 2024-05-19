package function 

import (
	"testing"
)

func TestFunction(t *testing.T) {
	expected :=  2 
	result := BinToInt("10")
	if expected != result {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}