package main

import (
	"fmt"
	"github.com/blang/semver/v4"
	"os/exec"
	"strings"
)

func main() {
	currentVersion := findCurrentVersion()
	fmt.Printf("Current version: %s\n", currentVersion)

	nextVersion, done := nextVersion(currentVersion)
	if done {
		return
	}
	fmt.Printf("Next version: %s\n", nextVersion)

	gitTag(nextVersion)

	gitPushTags()
}

func gitPushTags() string {
	return execute("git", "push", "--tags")
}

func gitTag(nextTag string) string {
	return execute("git", "tag", nextTag)
}

func nextVersion(tag string) (string, bool) {
	version, err := semver.Make(tag[1:])
	if err != nil {
		fmt.Println(err.Error())
		return "", true
	}

	_ = version.IncrementMinor()
	nextVersion := fmt.Sprintf("v%s", version)
	return nextVersion, false
}

func findCurrentVersion() string {
	tag := execute("git", "describe", "--tags", "--match", "v*", "--abbrev=0")
	return tag
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
