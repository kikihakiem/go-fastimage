package fastimage

import (
	"testing"
)

func TestBMPImage(t *testing.T) {
	t.Parallel()

	testCases := map[string]ImageSize{
		"http://www.ac-grenoble.fr/ien.vienne1-2/spip/IMG/bmp_Image004.bmp": ImageSize{477, 358},
		"http://eeweb.poly.edu/~yao/EL5123/image/lena_gray.bmp": ImageSize{512, 512},
		"https://samples.libav.org/image-samples/money-24bit-os2.bmp": ImageSize{455, 341},
		"https://samples.libav.org/image-samples/bmp-files/test4os2v2.bmp": ImageSize{300, 22},
		"http://cd.textfiles.com/monstmedia/IMAGES/JUR_OS2.BMP": ImageSize{640, 400},
	}

	for url, expectedSize := range testCases {
		imagetype, size, err := DetectImageType(url)
		if err != nil {
			t.Error("Failed to detect image type")
		}

		if imagetype != BMP {
			t.Error("Image is not BMP")
		}

		if size.Width != expectedSize.Width {
			t.Errorf("Image width is wrong. Expected %d, got %d", expectedSize.Width, size.Width)
		}

		if size.Height != expectedSize.Height {
			t.Errorf("Image height is wrong. Expected %d, got %d", expectedSize.Height, size.Height)
		}
	}
}
