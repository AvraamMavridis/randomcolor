package randomcolor

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// RGBColor RBG Color Type
type RGBColor struct {
	red   int
	green int
	blue  int
}

// HSVColor HSV Color Type
type HSVColor struct {
	hue        int
	saturation int
	value      int
}

// GetHex Converts a decimal number to hex representations
func GetHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

// GetRandomColorInRgb Returns a random RGBColor
func GetRandomColorInRgb() RGBColor {
	rand.Seed(time.Now().UnixNano())
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	c := RGBColor{red, green, blue}
	return c
}

// GetRandomColorInHex returns a random color in HEX format
func GetRandomColorInHex() string {
	color := GetRandomColorInRgb()
	hex := "#" + GetHex(color.red) + GetHex(color.green) + GetHex(color.blue)
	return hex
}

// GetRandomColorInHSV returns a random color in HSV format
func GetRandomColorInHSV() HSVColor {
	color := GetRandomColorInRgb()
	math.Max(float64(color.green), float64(color.blue))

	return HSVColor{0, 0, 0}
}
