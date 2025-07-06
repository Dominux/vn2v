package main

import (
	vnextractor "github.com/Dominux/vn2vn/vn_extractor"
	"github.com/google/uuid"
)

const basePath string = "/tmp/vn2vn"
const idStr string = "6d0d5ffe-cccc-4268-8858-6518c0fa85ca"

func main() {
	vnExtractor, err := vnextractor.NewVNExtractor(basePath)
	if err != nil {
		panic(err)
	}

	id := uuid.MustParse(idStr)
	if err := vnExtractor.ExtractAudio(id); err != nil {
		panic(err)
	}
}
