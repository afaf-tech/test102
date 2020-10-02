package main

import "testing"

func TestCount(t *testing.T) {
	if countStringNumber("5 + 2 - 1 - 2") != 4 {
		t.Error("Expected '5 + 2 - 1 - 2' equals to 4")
	}
	if countStringNumber("-7 + 5") != -2 {
		t.Error("Expected '-7 + 5' equals to -2")
	}
}
