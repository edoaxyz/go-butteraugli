#pragma once
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif
  typedef void* ImageInterface;
  ImageInterface ImageInterfaceInit(char *data, int width, int height);
  double Compare(ImageInterface im1, ImageInterface im2);
  void FreeImage(ImageInterface im);
#ifdef __cplusplus
}
#endif
