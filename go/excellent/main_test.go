package main

import (
	"testing"
)

// Testing the EvenOrOdd function
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("EvenOrOdd(10) = %s; want even", result)
	}
}
