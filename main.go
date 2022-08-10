package main

import (
	"fmt"
	"gocr/pkg"
)

func main() {
	client := pkg.NewClient()

	client.SetImage("summoner-test.png")
	client.ProcessGrayscale(false)
	client.ProcessThreshold(true, 245)

	output := client.Text()

	fmt.Printf("%s", output)
}
