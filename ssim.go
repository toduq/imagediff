package main

import (
	"math"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// https://en.wikipedia.org/wiki/Structural_similarity

const (
	WINDOW_SIZE = 8
	K1          = 0.01
	K2          = 0.03
	C1          = 0.0001 // K1**2 * L (L = 1.0)
	C2          = 0.0009 // K2**2 * L
)

func Ssim(from, to *imagick.MagickWand) float64 {
	width := int(from.GetImageWidth())
	height := int(from.GetImageHeight())
	cells := (width / WINDOW_SIZE) * (height / WINDOW_SIZE)
	colorSsims := [3]float64{0.0, 0.0, 0.0}
	for i := 0; i < cells; i++ {
		x := (i % (width / WINDOW_SIZE)) * WINDOW_SIZE
		y := (i / (height / WINDOW_SIZE)) * WINDOW_SIZE

		// collect pixel color value
		pixels := [3][2][]float64{{{}, {}}, {{}, {}}, {{}, {}}} // [r,g,b][from,to][0-63]
		for dy := 0; dy < WINDOW_SIZE; dy++ {
			for dx := 0; dx < WINDOW_SIZE; dx++ {
				f, _ := from.GetImagePixelColor(x, y)
				t, _ := to.GetImagePixelColor(x, y)
				pixels[0][0] = append(pixels[0][0], f.GetRed())
				pixels[1][0] = append(pixels[1][0], f.GetGreen())
				pixels[2][0] = append(pixels[2][0], f.GetBlue())
				pixels[0][1] = append(pixels[0][1], t.GetRed())
				pixels[1][1] = append(pixels[1][1], t.GetGreen())
				pixels[2][1] = append(pixels[2][1], t.GetBlue())
				f.Destroy()
				t.Destroy()
			}
		}
		// calculate ssim on its window
		for color := 0; color < 3; color++ {
			fromAve := vectorAverage(pixels[color][0])
			toAve := vectorAverage(pixels[color][1])
			fromDist := vectorDistort(pixels[color][0], fromAve)
			toDist := vectorDistort(pixels[color][1], toAve)
			covariance := vectorCovariance(pixels[color][0], pixels[color][1], fromAve, toAve)
			ssimUpper := (2*fromAve*toAve + C1) * (2*covariance + C2)
			ssimLower := (math.Pow(fromAve, 2) + math.Pow(toAve, 2) + C1) * (fromDist + toDist + C2)
			ssim := ssimUpper / ssimLower
			colorSsims[color] += ssim
		}
	}
	return (colorSsims[0] + colorSsims[1] + colorSsims[2]) / float64(3*cells)
}

func vectorAverage(vec []float64) float64 {
	sum := 0.0
	for _, v := range vec {
		sum += v
	}
	return sum / float64(len(vec))
}

func vectorDistort(vec []float64, ave float64) float64 {
	sum := 0.0
	for _, v := range vec {
		sum += math.Pow(v-ave, 2)
	}
	return sum / float64(len(vec))
}

func vectorCovariance(vec1, vec2 []float64, ave1, ave2 float64) float64 {
	sum := 0.0
	for i := 0; i < len(vec1); i++ {
		sum += (vec1[i] - ave1) * (vec2[i] - ave2)
	}
	return sum / float64(len(vec1))
}
