package vnextractor

type VNExtractor struct {
	basePath *string
}

func NewVNExtractor(basePath *string) *VNExtractor {
	return &VNExtractor{basePath: basePath}
}

func ExtractAudio(pathToVN *string) {
	// fullPath :=
	command := "ffmpeg -i "
}
