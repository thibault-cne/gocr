package pkg

import (
	"fmt"
	"gocr/pkg/effects"
	"image"
	"image/png"
	"os"
)

func (c *Client) saveProcessedImage(img image.Image) {
	f, err := os.Create("./out/out_" + c.ImageName)

	if err != nil {
		fmt.Printf("An error occured while creating out picture : %s", err.Error())
	}
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		fmt.Printf("An error occured while encoding picture : %s", err.Error())
	}
}

func (c *Client) retrieveProcessedImage() *image.Image {
	f, err := os.Open("./out/out_" + c.ImageName)

	if err != nil {
		fmt.Printf("An error occured on image open : %s\n", err.Error())
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	if err != nil {
		fmt.Printf("An error occured on image decode : %s\n", err.Error())
	}

	return &image
}

func (c *Client) ProcessThreshold(onOutput bool, level uint8) {
	var img *image.Gray

	if onOutput {
		cImg := c.retrieveProcessedImage()
		img = effects.Threshold(*cImg, level)
	} else {
		img = effects.Threshold(*c.Image, level)
	}

	c.saveProcessedImage(img)
}

func (c *Client) ProcessGrayscale(onOutput bool, config ...float64) {
	var img *image.RGBA

	if onOutput {
		cImg := c.retrieveProcessedImage()
		img = effects.Grayscale(*cImg, config...)
	} else {
		img = effects.Grayscale(*c.Image, config...)
	}

	c.saveProcessedImage(img)
}
