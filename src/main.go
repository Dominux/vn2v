package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// lmao.PrintSht()
	_, err := exec.Command("echo", "ebal mat' tvou suka!").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}

	// output := string(cmd)
	// println(output)
}
