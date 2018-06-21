package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
  "go-perlin"
  "math/rand"
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 100
)




var (
  f uint8=0
  l uint8=255
	blue color.Color = color.RGBA{f, f, l, l}
	black color.Color = color.RGBA{f, f, f, l}
	white  color.Color = color.RGBA{l, l, l, l}
)

func main() {

  p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	m := image.NewRGBA(image.Rect(0, 0, 640, 480)) //*NRGBA (image.Image interface)

  draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
  
    var j int
    j=0
    for x := 0.; x < 640; x++ {
      var k int
      k=0
      for y := 0.; y < 480; y++ {
        var r float64
        r=p.Noise2D(x/10, y/10)
        if r < 0.0 {
          m.Set(j, k, black)
        } else {
          m.Set(j, k, white)
        }
        k++
      }
      j++
    }

	w, _ := os.Create("perlin.png")
	defer w.Close()
	png.Encode(w, m)
}