package main

import (
	"github.com/google/uuid"
)

const basePath string = "/home/dominux/Downloads"

// const basePath string = "/tmp/vn2vn"
const idStr string = "6d0d5ffe-cccc-4268-8858-6518c0fa85ca"

func main() {
	vnExtractor, err := NewVNExtractor(basePath)
	if err != nil {
		panic(err)
	}

	id := uuid.MustParse(idStr)
	if err := vnExtractor.ExtractAudio(id); err != nil {
		panic(err)
	}
}
