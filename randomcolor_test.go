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
	if hsv.hue != expected {
		t.Errorf("Hue was not calculated correctly, got: %f, expected: %f.", hsv.hue, expected)
	}

	expected2 := 23.076923076923077
	if hsv.saturation != expected2 {
		t.Errorf("Saturation was not calculated correctly, got: %f, expected: %f.", hsv.saturation, expected2)
	}

	expected3 := 50.98039215686274
	if hsv.value != expected3 {
		t.Errorf("Saturation was not calculated correctly, got: %f, expected: %f.", hsv.value, expected3)
	}
}
