#include <butteraugli/butteraugli.cc>
#include "go-butteraugli.h"

using namespace butteraugli;

// Turns an interleaved RGBA buffer into 4 planes for each color channel
void planarize(std::vector<ImageF> &img,
               const uint8_t *rgba,
               int width,
               int height,
               float gamma = 2.2)
{
    assert(img.size() == 0);
    img.push_back(ImageF(width, height));
    img.push_back(ImageF(width, height));
    img.push_back(ImageF(width, height));
    img.push_back(ImageF(width, height));
    for (int y = 0; y < height; y++)
    {
        float *const row_r = img[0].Row(y);
        float *const row_g = img[1].Row(y);
        float *const row_b = img[2].Row(y);
        float *const row_a = img[3].Row(y);
        for (int x = 0; x < width; x++)
        {
            row_r[x] = 255.0 * pow(rgba[(y * width + x) * 4 + 0] / 255.0, gamma);
            row_g[x] = 255.0 * pow(rgba[(y * width + x) * 4 + 1] / 255.0, gamma);
            row_b[x] = 255.0 * pow(rgba[(y * width + x) * 4 + 2] / 255.0, gamma);
            row_a[x] = 255.0 * pow(rgba[(y * width + x) * 4 + 3] / 255.0, gamma);
        }
    }
}

ImageInterface ImageInterfaceInit(char *data, int width, int height)
{
    std::vector<ImageF> *image = new std::vector<ImageF>();
    planarize(*image, (uint8_t*)data, width, height);
    return (void *)(image);
}

double Compare(ImageInterface im1, ImageInterface im2)
{
    ImageF diffmap;
    double diffvalue;
    if (!ButteraugliInterface(*((std::vector<ImageF> *)im1), *((std::vector<ImageF> *)im2), 1.0, diffmap, diffvalue))
    {
        return -1;
    }
    return diffvalue;
}

void FreeImage(ImageInterface im) {
    ((std::vector<ImageF> *)im)->clear();
    delete ((std::vector<ImageF> *)im);
}