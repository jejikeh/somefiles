package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"os"
	"strings"
)

var genDirsCount = 0
var genFilesCount = 0
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	args := os.Args

	path := args[1]

	if strings.HasPrefix(path, "-") {
		path = "."
	}

	i := flag.Int("it", 1, "iterations")
	c := flag.Float64("chance", 0.5, "chance to generate a dir")
	d := flag.Bool("dir", true, "generate dirs")

	flag.Parse()

	genRandomDirectory(path, *i, 1.0-*c, *d)

	fmt.Printf("Generate dirs : %d\n", genDirsCount)
	fmt.Printf("Generate files: %d\n", genFilesCount)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func genRandomDirectory(dir string, it int, chance float64, genDir bool) string {
	dirName := dir + "/" + RandStringRunes(rand.Intn(10)+2)
	os.Mkdir(dirName, os.ModePerm)

	for i := 0; i < it+rand.Intn(it+1); i++ {
		if rand.Float64() > chance && genDir {
			genDirsCount++
			genRandomDirectory(dirName, rand.Intn(it+1), chance, genDir)
		}

		fileName := dirName + "/" + RandStringRunes(rand.Intn(10)+2)

		// @Cleanup: Move this to separate function

		if rand.Intn(2) == 0 {
			go genRandomImage(&fileName)
		} else {
			go genRandomText(&fileName)
		}

		fmt.Printf("%d file generated: %s\n", genFilesCount, fileName)
	}

	return dirName
}

func genRandomText(filename *string) {
	*filename += ".txt"

	f, _ := os.Create(*filename)
	defer f.Close()

	genFilesCount++

	f.WriteString(RandStringRunes(rand.Intn(1000)))
}

func genRandomImage(fileName *string) {
	*fileName += ".jpg"

	f, _ := os.Create(*fileName)
	defer f.Close()

	genFilesCount++

	width, height := rand.Intn(640), rand.Intn(480)
	background := color.RGBA{uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256))}

	img := createImage(width, height, background)
	for i := 0; i < width*height; i++ {
		img.SetRGBA(i%width, i/width, color.RGBA{uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256)), uint8(rand.Int31n(256))})
	}

	var opt jpeg.Options
	opt.Quality = rand.Intn(100)
	jpeg.Encode(f, img, &opt)
}

func createImage(width int, height int, background color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)

	img := image.NewRGBA(rect)

	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.Point{}, draw.Src)

	return img
}
