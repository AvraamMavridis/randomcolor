package randomcolor

import "testing"

func TestHex(t *testing.T) {
	hex := GetHex(10)
	expected := "0a"
	if hex != expected {
		t.Errorf("Hex was convertion was incorrect, got: %s, want: %s.", hex, expected)
	}

	hex2 := GetHex(255)
	expected2 := "ff"
	if hex2 != expected2 {
		t.Errorf("Hex convertion was incorrect, got: %s, want: %s.", hex2, expected2)
	}
}
