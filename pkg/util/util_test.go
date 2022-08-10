package util

import (
	"image"
	"image/color"
	"testing"
)

func TestRank(t *testing.T) {
	cases := []struct {
		value    color.RGBA
		expected float64
	}{
		{
			value:    color.RGBA{0, 0, 0, 0},
			expected: 0,
		},
		{
			value:    color.RGBA{255, 255, 255, 255},
			expected: 255,
		},
		{
			value:    color.RGBA{130, 130, 130, 80},
			expected: 130,
		},
		{
			value:    color.RGBA{70, 230, 55, 255},
			expected: 164.5,
		},
	}

	for _, c := range cases {
		actual := Rank(c.value)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v", "Rank", c.expected, actual)
		}
	}
}

func TestRGBAImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.RGBA
		b        *image.RGBA
		expected bool
	}{
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 3),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: true,
		},
		{
			a: &image.RGBA{},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a:        &image.RGBA{},
			b:        &image.RGBA{},
			expected: true,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := RGBAImageEqual(c.a, c.b)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v. Images a : %#v b : %#v", "RGBAImageEqual", c.expected, actual, c.a, c.b)
		}
	}
}

func TestGrayImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.Gray
		b        *image.Gray
		expected bool
	}{
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a:        &image.Gray{},
			b:        &image.Gray{},
			expected: true,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b:        &image.Gray{},
			expected: false,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := GrayImageEqual(c.a, c.b)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v. Images a : %#v b : %#v", "GrayImageEqual", c.expected, actual, c.a, c.b)
		}
	}
}
