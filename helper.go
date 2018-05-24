package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// return the current dir where you are to run the binary
func getPWD() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// return the dir of binary no matter where you run it
func getBinaryPath() (string, error) {
	dir, err := os.Executable()
	return dir, err
}

// this is similar to getBianryPath that return bianry dir
func absBinary() (string, error) {
	dir, err := filepath.Abs(path.Join(os.Args[0]))
	return dir, err
}

// match func check if a file match this pattern
func match(fileName string) bool {
	r := regexp.MustCompile("\\s\\(\\d+\\sof\\s\\d+\\)")
	return r.MatchString(fileName)
}

func rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func formNewName(oldName string, seq int) string {
	r := regexp.MustCompile("\\s\\(\\d+\\sof\\s\\d+\\)")
	suffix := fmt.Sprintf("_%v", seq)
	replaced := r.ReplaceAllString(oldName, suffix)
	return replaced
}
