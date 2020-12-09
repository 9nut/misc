// An exercise from Russ Cox's "A Tour of Go" course at USENIX 2010
// Here's the original paper: http://swtch.com/usenix/go-course.pdf

package main

import (
	"image"
	"image/color"
	"math"
)

// Newton attractor supports an image.Image interface
type Newton struct {
	Dx int        // Image width
	Dy int        // Image height
	Z0 complex128 // minimum starting point
	Dz complex128 // range of values to search
	N  int        // color
}

func NewNewton(dx, dy int, z0, dz complex128, n int) *Newton {
	return &Newton{dx, dy, z0, dz, n}
}

func (n *Newton) Bounds() image.Rectangle {
	return image.Rect(0, 0, n.Dx, n.Dy)
}

func (n *Newton) ColorModel() color.Model {
	return color.NRGBAModel
}

func (n *Newton) At(x, y int) color.Color {
	fx := float64(x) / float64(n.Dx)
	fy := float64(y) / float64(n.Dy)

	Z := n.Z0 + complex(real(n.Dz)*fx, imag(n.Dz)*fy)
	r := []complex128{complex(1, 0), complex(-1, 1), complex(-1, -1)}
	c := []int{0, 0, 0}
	for i, v := range r {
		z := Z
		for c[i] = n.N; c[i] > 0; c[i]-- {
			z0 := z
			z -= (z*z*z - v) / (3 * z * z)
			if math.Abs(real(z)-real(z0)) < 1e-7 {
				break
			}
		}
	}

	var shift uint
	for i := n.N - 1; i > 0; i >>= 1 {
		shift++
	}
	adj := func(a int) uint8 {
		foo := uint8(a & 1)
		return uint8((a<<8)>>shift) | foo
		// or use return uint8((a*255)/n.N)
	}
	return color.NRGBA{adj(c[2]), adj(c[0]), adj(c[1]), 0xff}
}
