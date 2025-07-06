package vnextractor

import (
	"fmt"
	"os"

	"github.com/Dominux/vn2vn/constants"
	"github.com/google/uuid"
)

type VNExtractor struct {
	basePath string
}

func NewVNExtractor(basePath string) (*VNExtractor, error) {
	// Removing base path
	removeBasePath(basePath)

	// Making base path
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &VNExtractor{basePath: basePath}, nil
}

func (v *VNExtractor) ExtractAudio(id uuid.UUID) {
	fullPath := v.buildFullPath(id, constants.InputVN)
	command := fmt.Sprintf("ffmpeg -i %s")
}

func (v *VNExtractor) buildFullPath(id uuid.UUID, filename string) string {
	return fmt.Sprintf("%s/%s/%s", v.basePath, id.String(), filename)
}

func removeBasePath(basePath string) error {
	return os.RemoveAll(basePath)
}
