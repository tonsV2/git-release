package main

import (
	"flag"
	"fmt"
	"github.com/blang/semver/v4"
	"os/exec"
	"strings"
)

func main() {
	var strategy string
	flag.StringVar(&strategy, "strategy", "minor", "Bump strategy, can be either major, minor or patch")
	var dryRun bool
	flag.BoolVar(&dryRun, "dry-run", false, "Dry run? Don't execute commands")
	// safe run? Prompt before pushing
	flag.Parse()

	currentVersion := findCurrentVersion()
	fmt.Printf("Current version: %s\n", currentVersion)

	nextVersion, done := nextVersion(currentVersion, strategy)
	if done {
		return
	}
	fmt.Printf("Next version: %s\n", nextVersion)

	gitTag(nextVersion, dryRun)

	gitPushTags(dryRun)
}

func gitPushTags(dryRun bool) string {
	if dryRun {
		return ""
	}
	return execute("git push --tags")
}

func gitTag(nextTag string, dryRun bool) string {
	if dryRun {
		return ""
	}
	return execute(fmt.Sprintf("git tag %s", nextTag))
}

func nextVersion(tag string, strategy string) (string, bool) {
	version, err := semver.Make(tag[1:])
	if err != nil {
		fmt.Println(err.Error())
		return "", true
	}

	switch strategy {
	case "major":
		_ = version.IncrementMajor()
	case "minor":
		_ = version.IncrementMinor()
	case "patch":
		_ = version.IncrementPatch()
	}

	nextVersion := fmt.Sprintf("v%s", version)
	return nextVersion, false
}

func findCurrentVersion() string {
	tag := execute("git describe --tags --match v* --abbrev=0")
	return tag
}

func execute(command string) string {
	fmt.Println(command)
	fields := strings.Fields(command)
	app := fields[0]

	cmd := exec.Command(app, fields[1:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return strings.TrimSpace(string(stdout))
}
