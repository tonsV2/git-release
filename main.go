package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	tag := execute("git", "describe", "--tags", "--match", "v*", "--abbrev=0")
	fmt.Println(tag)
}

func execute(param ...string) string {
	app := param[0]

	cmd := exec.Command(app, param[1:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return strings.TrimSpace(string(stdout))
}
