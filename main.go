package main

import (
	"fmt"
	"github.com/blang/semver/v4"
	"os/exec"
	"strings"
)

func main() {
	tag := execute("git", "describe", "--tags", "--match", "v*", "--abbrev=0")
	fmt.Printf("Current version: %s\n", tag)

	version, err := semver.Make(tag[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_ = version.IncrementMinor()
	nextTag := fmt.Sprintf("v%s", version)

	fmt.Printf("Next version: %s\n", nextTag)
}

func execute(param ...string) string {
	fmt.Println(strings.Join(param, " "))

	app := param[0]

	cmd := exec.Command(app, param[1:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return strings.TrimSpace(string(stdout))
}
