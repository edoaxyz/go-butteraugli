package butteraugli_test

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/edoaxyz/go-butteraugli"
)

var paths = []string{
	"images/quality_100.jpg",
	"images/quality_90.jpg",
	"images/quality_80.jpg",
	"images/quality_70.jpg",
	"images/quality_60.jpg",
}

func TestButteraugli(t *testing.T) {
	last_result := -0.1

	fqOriginal, _ := os.Open("images/quality_100.jpg")
	imgOriginal, _ := jpeg.Decode(fqOriginal)
	butterOriginal := butteraugli.BuildButteraugliImage(imgOriginal)

	for _, path := range paths {
		fqImg, _ := os.Open(path)
		img, _ := jpeg.Decode(fqImg)
		butterImg := butteraugli.BuildButteraugliImage(img)

		compare := butterOriginal.Compare(butterImg)
		if compare < last_result {
			t.Errorf("Got distance <= last computed distance while comparing %s with %s", fqOriginal.Name(), fqImg.Name())
		}
		last_result = compare

		defer butterImg.Close()
		defer fqImg.Close()
	}

	defer fqOriginal.Close()
	defer butterOriginal.Close()
}
