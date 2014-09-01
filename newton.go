// An exercise from Russ Cox's "A Tour of Go" course at USENIX 2010
// Here's the original paper: http://swtch.com/usenix/go-course.pdf

package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

type Newton struct {
	Dx, Dy int
	Z0, Dz complex128
	N      int
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
	// log.Println("Newton:At")
	fx := float64(x) / float64(n.Dx)
	fy := float64(y) / float64(n.Dy)

	Z := n.Z0 + complex(real(n.Dz)*fx, imag(n.Dz)*fy)
	r := []complex128{complex(1, 0), complex(-1, 1), complex(-1, -1)}
	c := []int{0, 0, 0}
	for i, v := range r {
		z := Z
		// log.Println(i, v)
		for c[i] = n.N; c[i] > 0; c[i]-- {
			z0 := z
			z -= (z*z*z - v) / (3 * z * z)
			if cmplx.Abs(z-z0) < 1e-7 {
				break
			}
			// log.Println(i, c[i], z-z0)
		}
	}

	// log.Println(c)
	adj := func(a int) uint8 { return uint8((0xff * a) / n.N) }
	return color.NRGBA{adj(c[2]), adj(c[0]), adj(c[1]), 0xff}
}

func main() {
	var z0r *float64 = flag.Float64("z0r", -3, "z0 real part")
	var z0i *float64 = flag.Float64("z0i", -3, "z0 imaginary part")
	var z1r *float64 = flag.Float64("z1r", 3, "z1 real part")
	var z1i *float64 = flag.Float64("z1i", 3, "z1 imaginary part")

	flag.Parse()

	z0 := complex(*z0r, *z0i)
	z1 := complex(*z1r, *z1i)
	dz := z1 - z0
	n := NewNewton(2048, 2048, z0, dz, 64)
	name := fmt.Sprintf("newton%+f%+fi%+f%+fi.png", real(z0), imag(z0), real(z1), imag(z1))
	if store, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666); err == nil {
		defer store.Close()
		png.Encode(store, n)
	} else {
		log.Println("newton: ", err)
	}
}
