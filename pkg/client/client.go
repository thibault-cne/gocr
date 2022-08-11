package client

import (
	"fmt"
	"image"
	"os"
	"os/exec"
	"path/filepath"
)

type Client struct {
	// All languages to be detected. By default it is set to english only.
	Languages []string

	// Can indicate the path to tessdata folder.
	TessDataPrefix string

	// Save the image to analyse.
	Image image.Image

	// Save the image name.
	ImageName string

	// Save the image path
	ImagePath string
}

func NewClient() *Client {
	// By default we create a `out` folder for the processed images.
	mkdirCommand := exec.Command("mkdir", "-p", "out")
	mkdirCommand.Start()

	return &Client{
		Languages: []string{"eng"},
	}
}

func (client *Client) SetLanguages(langs ...string) {
	if len(langs) == 0 {
		fmt.Printf("Languages list cannot be empty.")
		return
	}

	client.Languages = langs
}

func (client *Client) SetImage(imagePath string) {
	if imagePath == "" {
		fmt.Printf("File path cannot be empty.")
		return
	}

	f, err := os.Open(imagePath)

	if err != nil {
		fmt.Printf("An error occured image open : %s\n", err.Error())
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	if err != nil {
		fmt.Printf("An error occured image decode : %s\n", err.Error())
	}

	client.Image = image
	client.ImageName = filepath.Base(imagePath)
	client.ImagePath = imagePath
}

// Runs the tesseract command with the image of the client
// The args parameter is used to run specific arguments in the tesseract command.
// By default it runs `stdout -l eng`.
func (c *Client) Text(args ...string) string {
	_, err := exec.LookPath("tesseract")

	if err != nil {
		fmt.Printf("Error - tesseract not found. You need to download tesseract to use this package.")
		return ""
	}

	var (
		cmd     *exec.Cmd
		imgPath string
	)

	imgPath = c.ImagePath

	if len(args) > 0 {
		tempArgs := []string{imgPath}
		tempArgs = append(tempArgs, args...)
		cmd = exec.Command("tesseract", tempArgs...)
	} else {
		tempArgs := []string{imgPath, "stdout", "-l"}
		tempArgs = append(tempArgs, c.Languages...)
		cmd = exec.Command("tesseract", tempArgs...)
	}

	stdout := startCommand(cmd)

	output := ""

	for stdout.Scan() {
		output += stdout.Text() + "\n"
	}

	return output
}
