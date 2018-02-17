package main

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// ListFiles returns list of files and their total size in startDir and its sub-directiries
func ListFiles(startDir string) (files []string, size int64) {
	filesAndDirs, err := ioutil.ReadDir(startDir)

	if err != nil {
		if strings.Contains(err.Error(), "Access is denied") {
			log.Println(err)
		} else {
			log.Panicln(err)
		}
	}

	files = make([]string, 0, 100)
	
	for _, f := range filesAndDirs {
		if f.IsDir() {
			fl, sz := ListFiles(path.Join(startDir, f.Name()))
			files = append(files, fl...)
			size += sz
		} else {
			files = append(files, path.Join(startDir, f.Name()))
			size += f.Size()
		}
	}

	return files, size
}

func main() {
	files, size := ListFiles("c:\\temp")
	log.Printf("Total number of files: %d, total size: %d", len(files), size)
}