# go-butteraugli

A Go wrapper for [butteraugli](https://github.com/google/butteraugli)

```go
package main

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/edoaxyz/go-butteraugli"
)

func main() {
	fqOriginal, _ := os.Open("original.jpg")
	imgOriginal, _ := jpeg.Decode(fqOriginal)
	butterOriginal := butteraugli.BuildButteraugliImage(imgOriginal)

    fqImg, _ := os.Open("optimized.jpg")
    img, _ := jpeg.Decode(fqImg)
    butterImg := butteraugli.BuildButteraugliImage(img)

    compare := butterOriginal.Compare(butterImg)

    defer butterImg.Close() // Remember to close the image to free memory!
    defer fqImg.Close()

	defer butterOriginal.Close() // Remember to close the image to free memory!
	defer fqOriginal.Close()
}

```