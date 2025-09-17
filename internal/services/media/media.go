package service

import (
	"fmt"
	media "txrnxp-whats-happening/external/media/files"
)

type MediaService struct {
	imageProvider media.MediaStorageProvider
}

func NewMediaService(imageProvider media.MediaStorageProvider) *MediaService {
	return &MediaService{
		imageProvider: imageProvider,
	}
}

// UploadMedia handles file uploads
func (ms *MediaService) UploadMedia(fileName string, data []byte) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("file data cannot be empty")
	}
	url, err := ms.imageProvider.UploadFile(fileName, data)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}
	return url, nil
}

// GetMediaURL builds the public URL
func (ms *MediaService) GetMediaURL(path string) string {
	return ms.imageProvider.RetrieveFile(path)
}
