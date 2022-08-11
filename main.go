package main

import (
	"fmt"

	"gocr/pkg/client"
)

func main() {
	client := client.NewClient()

	// We set a new image
	client.SetImage("summoner-test.png")

	// We add a grayscale effect to it
	client.Grayscale(true)
	// We add a threshold of level 240
	client.Threshold(true, 240)

	// We retrieve the output
	output := client.Text()

	fmt.Printf("%s", output)
}
