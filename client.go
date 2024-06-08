package bnrimage

import (
	"github.com/disintegration/imaging"
)

type Client interface {
	ResizeImage(imageName, fullPath, dstPath string) (*ImageData, error)
}

const (
	defaultImageWidth  = 800
	defaultImageHeight = 0
)

type defaultClient struct {
	rFilter            imaging.ResampleFilter
	defaultImageWidth  int
	defaultImageHeight int
}

func NewClient() Client {
	filter := imaging.Lanczos

	return &defaultClient{
		rFilter:            filter,
		defaultImageWidth:  defaultImageWidth,
		defaultImageHeight: defaultImageHeight,
	}
}
