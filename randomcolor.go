package randomcolor

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Color RBG Type
type Color struct {
	red   int
	green int
	blue  int
}

// GetHex Converts a decimal number to hex representations
func GetHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

// GetRandomColorInRgb Returns a random Color
func GetRandomColorInRgb() Color {
	rand.Seed(time.Now().UnixNano())
	var red = rand.Intn(255)
	var green = rand.Intn(255)
	var blue = rand.Intn(255)
	c := Color{red, green, blue}
	return c
}

// GetRandomColorInHex returns a random color in HEX format
func GetRandomColorInHex() string {
	color := GetRandomColorInRgb()
	hex := "#" + GetHex(color.red) + GetHex(color.green) + GetHex(color.blue)
	return hex
}

func GetRandomColorInHSV() {
	color := GetRandomColorInRgb()
	math.Max(float64(color.green), float64(color.blue))
}
