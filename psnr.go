package main

import (
	"math"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// https://ja.wikipedia.org/wiki/%E3%83%94%E3%83%BC%E3%82%AF%E4%BF%A1%E5%8F%B7%E5%AF%BE%E9%9B%91%E9%9F%B3%E6%AF%94

func Psnr(from, to *imagick.MagickWand) float64 {
	width := int(from.GetImageWidth())
	height := int(from.GetImageHeight())
	colorMse := [3]float64{}
	for x := 0; x < width; x++ {
		for y := 0; y < 50; y++ {
			f, _ := from.GetImagePixelColor(x, y)
			t, _ := to.GetImagePixelColor(x, y)
			colorMse[0] += math.Pow(f.GetRed()-t.GetRed(), 2.0)
			colorMse[1] += math.Pow(f.GetGreen()-t.GetGreen(), 2.0)
			colorMse[2] += math.Pow(f.GetBlue()-t.GetBlue(), 2.0)
			f.Destroy()
			t.Destroy()
		}
	}
	mse := 0.0
	for i := 0; i < 3; i++ {
		mse += colorMse[i] / float64(width*height)
	}
	mse /= 3
	if mse == 0.0 {
		return 0.0
	}
	psnr := -10 * math.Log10(mse)
	return psnr
}
