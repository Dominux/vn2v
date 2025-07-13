package main

import (
	"fmt"
	"os"
	"os/exec"
)

type VNExtractor struct {
	basePath string
}

func NewVNExtractor(basePath string) (*VNExtractor, error) {
	// Removing base path
	// if err := removeBasePath(basePath); err != nil {
	// 	return nil, err
	// }

	return &VNExtractor{basePath: basePath}, nil
}

func (v *VNExtractor) ExtractAudio() error {
	inputFullPath := v.buildFullPath(InputVN)
	outputFullPath := v.buildFullPath(InputAudio)
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i",
		inputFullPath,
		outputFullPath,
	)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (v *VNExtractor) ExtractImage() error {
	inputFullPath := v.buildFullPath(InputVN)
	outputFullPath := v.buildFullPath(Screenshot)
	cmd := exec.Command(
		"ffmpeg",
		"-ss",
		"00:00:02",
		"-y",
		"-i",
		inputFullPath,
		"-frames:v",
		"1",
		outputFullPath,
	)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (v *VNExtractor) buildFullPath(filename string) string {
	return fmt.Sprintf("%s%s", v.basePath, filename)
}

func removeBasePath(basePath string) error {
	if _, err := os.Stat(basePath); err != nil {
		return nil
	}
	return os.RemoveAll(basePath)
}
