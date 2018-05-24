package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	// "io/ioutil"
	"path/filepath"
)

var dir = flag.String("root", "", "root directory to traverse")
var stopDir = flag.String("stop", "", "stop dir. If empty, will travese all sub-directories")
var depth = flag.Int("depth", -1, "traverse depth. negative value will traverse all sub-directories.")

func main() {
	flag.Parse()
	current := ""
	if *dir == "" {
		fmt.Println("root is empty. start traversing from current dir")
		var err error
		current, err = getPWD()
		must(err)
	} else {
		current = *dir
	}
	// fi, _ := ioutil.ReadDir(current)
	// for _, f := range fi {
	// 	fmt.Println(f.Name(), f.IsDir())
	// }
	fmt.Println()
	fmt.Println("start walking")
	count := 0
	err := filepath.Walk(current, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		if match(fileName) {
			fmt.Println("find file: ", info.Name(), "with path: ", path)
			newName := formNewName(fileName, count)
			newPath := strings.Replace(path, fileName, newName, -1)
			err := rename(path, newPath)
			must(err)
			count++
		}
		return nil
	})
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
