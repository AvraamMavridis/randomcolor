package randomcolor

import (
	"testing"
)

func TestHex(t *testing.T) {
	hex := getHex(10)
	expected := "0a"
	if hex != expected {
		t.Errorf("Hex was convertion was incorrect, got: %s, want: %s.", hex, expected)
	}

	hex2 := getHex(255)
	expected2 := "ff"
	if hex2 != expected2 {
		t.Errorf("Hex convertion was incorrect, got: %s, want: %s.", hex2, expected2)
	}
}

func TestRGBtoHex(t *testing.T) {
	c := RGBColor{100, 130, 100}
	hsv := rgpToHSV(c)

	expected := 120.0
	if hsv.Hue != expected {
		t.Errorf("Hue was not calculated correctly, got: %f, expected: %f.", hsv.Hue, expected)
	}

	expected2 := 23.076923076923077
	if hsv.Saturation != expected2 {
		t.Errorf("Saturation was not calculated correctly, got: %f, expected: %f.", hsv.Saturation, expected2)
	}

	expected3 := 50.98039215686274
	if hsv.Value != expected3 {
		t.Errorf("Saturation was not calculated correctly, got: %f, expected: %f.", hsv.Value, expected3)
	}
}
