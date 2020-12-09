// An exercise from Russ Cox's "A Tour of Go" course at USENIX 2010
// Here's the original paper: http://swtch.com/usenix/go-course.pdf

package main

import (
	"flag"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	z0r := flag.Float64("z0r", -3, "z0 real part")
	z0i := flag.Float64("z0i", -3, "z0 imaginary part")
	z1r := flag.Float64("z1r", 3, "z1 real part")
	z1i := flag.Float64("z1i", 3, "z1 imaginary part")
	N := flag.Int("N", 32, "iterations")

	flag.Parse()

	// calculate the size
	// for now we normalize the size of the image to 2048 by (2048 * dy/dx)
	dx := *z1r - *z0r
	dy := *z1i - *z0i
	ratio := math.Abs(dy / dx)
	Dx := 2048
	Dy := int(2048 * ratio)
	log.Printf("Dx = %d, Dy = %d\n", Dx, Dy)

	z0 := complex(*z0r, *z0i)
	z1 := complex(*z1r, *z1i)
	dz := z1 - z0

	n := NewNewton(Dx, Dy, z0, dz, *N)
	png.Encode(os.Stdout, n)
}
