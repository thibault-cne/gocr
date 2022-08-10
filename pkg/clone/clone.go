package clone

import (
	"image"
	"image/draw"
)

func CloneAsRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}
