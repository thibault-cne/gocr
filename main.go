package main

import (
	"fmt"

	"gocr/pkg/client"
)

func main() {
	client := client.NewClient()

	client.SetImage("summoner-test.png")
	client.ProcessGrayscale(false)
	client.ProcessThreshold(true, 245)

	output := client.Text(true)

	fmt.Printf("%s", output)
}
