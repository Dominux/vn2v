package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/t2i", t2i)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func t2i(w http.ResponseWriter, r *http.Request) {
	inputPrompt := r.FormValue("prompt")

	cmd := exec.Command(
		"ollama",
		"-y",
		"-i",
		inputFullPath,
		outputFullPath,
	)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
}
