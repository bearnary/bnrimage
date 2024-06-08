package bnrimage

import (
	"strings"

	"github.com/disintegration/imaging"
)

func (c *defaultClient) ResizeImage(imageName, fullPath, dstPath string) (*ImageData, error) {

	src, err := imaging.Open(fullPath)
	if err != nil {
		return nil, err
	}

	dstImg := imaging.Resize(src, c.defaultImageWidth, c.defaultImageHeight, imaging.Lanczos)

	fs := strings.Split(fullPath, ".")
	ext := fs[len(fs)-1]

	newFileName := imageName + "." + ext
	dstFullPath := dstPath + "/" + newFileName

	err = imaging.Save(dstImg, dstFullPath)
	if err != nil {
		return nil, err
	}

	imgData := ImageData{
		Name:     newFileName,
		Path:     dstPath,
		FullPath: dstFullPath,
	}

	return &imgData, nil
}
