package main

import (
	"fmt"

	"gocr/pkg/client"
)

func main() {
	client := client.NewClient()

	client.SetImage("summoner-test.png")
	client.ProcessGrayscale(false, 0.3, 0.7, 0.1)
	client.ProcessThreshold(true, 245)

	output := client.Text(true)

	fmt.Printf("%s", output)
}
