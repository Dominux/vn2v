package main

const basePath string = "../data/"

func main() {
	vnExtractor, err := NewVNExtractor(basePath)
	if err != nil {
		panic(err)
	}

	if err := vnExtractor.ExtractAudio(); err != nil {
		panic(err)
	}

	if err := vnExtractor.ExtractImage(); err != nil {
		panic(err)
	}
}
