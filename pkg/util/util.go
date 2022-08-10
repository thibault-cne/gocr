package util

import (
	"image/color"
)

func Rank(color color.RGBA) float64 {
	return float64(color.R)*0.3 + float64(color.G)*0.6 + float64(color.B)*0.1
}
