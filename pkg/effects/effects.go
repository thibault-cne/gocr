package effects

import (
	"gocr/pkg/util"
	"image"
	"image/color"
	"image/draw"
)

func cloneAsRGBA(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}

func Grayscale(img image.Image, config ...float64) *image.RGBA {
	// Set the weight for the grayscale luma method. By default it respects the ITU-R recommendations
	rWeight := 0.2126
	gWeight := 0.7152
	bWeight := 0.0722

	// We verify that the sum of weight is equal to 1
	if len(config) == 3 && config[0]+config[1]+config[2] == 1 {
		rWeight = config[0]
		gWeight = config[1]
		bWeight = config[2]
	}

	src := cloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := uint8(float64(pix.R)*rWeight + float64(pix.G)*gWeight + float64(pix.B)*bWeight)
			dst.Set(x, y, color.RGBA{gray, gray, gray, pix.A})
		}
	}

	return dst
}

func Threshold(img image.Image, level uint8) *image.Gray {
	src := cloneAsRGBA(img)
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
