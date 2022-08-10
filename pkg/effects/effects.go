package effects

import (
	"gocr/pkg/clone"
	"gocr/pkg/util"
	"image"
	"image/color"
)

func Grayscale(img image.Image, config ...float64) *image.RGBA {
	// Set the weight for the grayscale luma method. By default it respects the ITU-R recommendations
	r := 0.2126
	g := 0.7152
	b := 0.0722

	// We verify that the sum of weight is equal to 1 else we apply the default values.
	if len(config) == 3 && float32(config[0])+float32(config[1])+float32(config[2]) == 1 {
		r = config[0]
		g = config[1]
		b = config[2]
	}

	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := uint8(float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5)
			dst.Set(x, y, color.RGBA{gray, gray, gray, pix.A})
		}
	}

	return dst
}

func Threshold(img image.Image, level uint8) *image.Gray {
	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewGray(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)

			r := util.Rank(pix)

			if uint8(r) >= level || pix.R == 0 && pix.G == 0 && pix.B == 0 && pix.A == 0 {
				dst.Set(x, y, color.White)
			} else {
				dst.Set(x, y, color.Black)
			}
		}
	}

	return dst
}
