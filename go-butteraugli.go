package butteraugli

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I./butteraugli
// #cgo windows LDFLAGS: -Wl,--allow-multiple-definition
// #cgo !windows LDFLAGS: -lm
// #include <stdlib.h>
// #include "go-butteraugli.h"
import "C"
import (
	"image"
	"unsafe"

	"golang.org/x/image/draw"
)

type ButteraugliImage struct {
	ptr unsafe.Pointer
}

func (im *ButteraugliImage) Close() error {
	defer C.FreeImage((C.ImageInterface)(im.ptr))
	return nil
}

func (im *ButteraugliImage) Compare(image ButteraugliImage) float64 {
	return (float64)(C.Compare((C.ImageInterface)(im.ptr), (C.ImageInterface)(image.ptr)))
}

func BuildButteraugliImage(src image.Image) ButteraugliImage {
	src_rgba := image.NewRGBA(src.Bounds())
	draw.Copy(src_rgba, image.Point{}, src, src.Bounds(), draw.Src, nil)

	str_data := C.CString(string(src_rgba.Pix))
	p := C.ImageInterfaceInit(str_data, C.int(src.Bounds().Dx()), C.int(src.Bounds().Dy()))
	defer C.free(unsafe.Pointer(str_data))

	return ButteraugliImage{unsafe.Pointer(p)}
}
