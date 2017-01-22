package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/gographics/imagick.v2/imagick"
)

var (
	method   string
	fromPath string
	toPath   string
)

func main() {
	flag.StringVar(&method, "m", "psnr", "Method [psnr, ssim]")
	flag.Parse()
	args := flag.Args()
	fromPath, toPath := args[0], args[1]

	imagick.Initialize()
	defer imagick.Terminate()

	from := loadImage(fromPath)
	to := loadImage(toPath)

	if from.GetImageWidth() != to.GetImageWidth() || from.GetImageHeight() != to.GetImageHeight() {
		panic("image size not match")
	}

	if method == "psnr" {
		fmt.Println(Psnr(from, to))
	} else {
		fmt.Println(Ssim(from, to))
	}
}

func loadImage(path string) *imagick.MagickWand {
	img := imagick.NewMagickWand()
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	err = img.ReadImage(path)
	if err != nil {
		panic(err)
	}
	return img
}
