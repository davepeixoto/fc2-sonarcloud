package main

import "testing"

func TestSAum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}
