package effects

import (
	"gocr/pkg/util"
	"image"
	"testing"
)

func TestGrayscale(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					130, 255, 255, 0x80,
					130, 255, 255, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					228, 228, 228, 0x80,
					228, 228, 228, 0x80,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Grayscale(c.value)

		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s : expected : %#v, actual : %#v", "Grayscale", c.expected, actual)
		}
	}
}

func TestGrayscaleWithConfig(t *testing.T) {
	cases := []struct {
		config   []float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			config: []float64{0.2345, 0.4354, 0.5554},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			config: []float64{0.2, 0.7, 0.1},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			config: []float64{0.2, 0.7, 0.1},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			config: []float64{0.3, 0.7, 0.1},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					130, 255, 255, 0x80,
					130, 255, 255, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					228, 228, 228, 0x80,
					228, 228, 228, 0x80,
				},
			},
		},
		{
			config: []float64{0.2, 0.7, 0.1},
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					130, 255, 255, 0x80,
					130, 255, 255, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					230, 230, 230, 0x80,
					230, 230, 230, 0x80,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Grayscale(c.value, c.config...)

		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s : expected : %#v, actual : %#v", "Grayscale", c.expected, actual)
		}
	}
}

func TestThreshold(t *testing.T) {
	cases := []struct {
		level    uint8
		value    image.Image
		expected *image.Gray
	}{
		{
			level: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
		},
		{
			level: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 4 * 3,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			level: 128,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 4 * 3,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0x00, 0xFF, 0x00,
					0xFF, 0x00, 0xFF,
				},
			},
		},
		{
			level: 255,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 4 * 3,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0x00, 0xFF, 0x00,
					0xFF, 0x00, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Threshold(c.value, c.level)

		if !util.GrayImageEqual(actual, c.expected) {
			t.Errorf("%s : with level %d expected : %#v, actual : %#v", "Threshold", c.level, c.expected, actual)
		}
	}
}
