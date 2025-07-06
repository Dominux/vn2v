package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type VNExtractor struct {
	basePath string
}

func NewVNExtractor(basePath string) (*VNExtractor, error) {
	// Removing base path
	// if err := removeBasePath(basePath); err != nil {
	// 	return nil, err
	// }

	// Making base path
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return nil, err
	}

	return &VNExtractor{basePath: basePath}, nil
}

func (v *VNExtractor) ExtractAudio(id uuid.UUID) error {
	inputFullPath := v.buildFullPath(id, InputVN)
	outputFullPath := v.buildFullPath(id, InputAudio)
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		inputFullPath,
		"-map",
		"0:a",
		"-acodec",
		"libmp3lame",
		outputFullPath,
	)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (v *VNExtractor) buildFullPath(id uuid.UUID, filename string) string {
	return fmt.Sprintf("%s/%s/%s", v.basePath, id.String(), filename)
}

func removeBasePath(basePath string) error {
	if _, err := os.Stat(basePath); err != nil {
		return nil
	}
	return os.RemoveAll(basePath)
}
