package validation

import (
	"errors"
	"strings"
)

func ValidationImages(imageName string, imageSize int) error {
	name := strings.ToLower(imageName)
	if !strings.HasSuffix(name, ".jpg") && !strings.HasSuffix(name, ".png") && !strings.HasSuffix(name, ".jpeg") {
		return errors.New("only supported file formats are jpg, jpeg and png")
	}

	if imageSize > 5003000 {
		return errors.New("image size cannot be more than 5 MB")
	}

	return nil
}


