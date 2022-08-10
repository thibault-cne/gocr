package main

import (
	"fmt"

	"gocr/pkg/client"
)

func main() {
	client := client.NewClient()

	client.SetImage("summoner-test.png")
	client.Grayscale(false, 0.3, 0.7, 0.1)
	client.Threshold(true, 245)

	output := client.Text(true)

	fmt.Printf("%s", output)
}
