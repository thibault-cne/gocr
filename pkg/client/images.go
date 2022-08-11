package client

import (
	"fmt"
	"gocr/pkg/effects"
	"image"
	"image/png"
	"os"
)

func (c *Client) saveProcessedImage(img image.Image, method string) {
	f, err := os.Create("./out/" + method + "_" + c.ImageName)

	if err != nil {
		fmt.Printf("An error occured while creating out picture : %s", err.Error())
	}
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		fmt.Printf("An error occured while encoding picture : %s", err.Error())
	}
}

// Process a threshold on the client image. You can save the processed threshold image in the client in
// by using the useAsNewImage boolean.
// You can set the threshold level with the level parameter.
func (c *Client) Threshold(useAsNewImage bool, level uint8) {
	img := effects.Threshold(c.Image, level)

	if useAsNewImage {
		c.Image = img
		c.ImagePath = "./out/threshold_" + c.ImageName
	}

	c.saveProcessedImage(img, "threshold")
}

// Process a grayscale on the client image. You can save the processed grayscale image in the client in
// by using the useAsNewImage boolean.
// You can set the grayscale weight with the config parameter `[]float64`. Make sure that the
// sum of the 3 floats is equal to 1.
func (c *Client) Grayscale(useAsNewImage bool, config ...float64) {
	img := effects.Grayscale(c.Image, config...)

	if useAsNewImage {
		c.Image = img
		c.ImagePath = "./out/grayscale_" + c.ImageName
	}

	c.saveProcessedImage(img, "grayscale")
}
