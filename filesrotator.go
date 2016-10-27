package filesrotator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

type sFiles []os.FileInfo

func (f sFiles) Len() int {
	return len(f)
}
func (f sFiles) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type sByFilename struct {
	sFiles
}

func (s sByFilename) Less(i, j int) bool {
	return s.sFiles[i].Name() > s.sFiles[j].Name()
}

type sByModTime struct {
	sFiles
}

func (s sByModTime) Less(i, j int) bool {
	return s.sFiles[i].ModTime().Unix() < s.sFiles[j].ModTime().Unix()
}

func RotateByFilename(directory string, maxfiles int) (err error) {
	directoryFiles, err := getDirFiles(directory)

	sort.Sort(sByFilename{directoryFiles})

	err = removeOlder(directory, directoryFiles, maxfiles)

	return
}

func getDirFiles(directory string) (files sFiles, err error) {
	var directoryFiles []os.FileInfo

	files, err = ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	for _, f := range files {
		if !f.IsDir() {
			directoryFiles = append(directoryFiles, f)
		}
	}

	return
}

func removeOlder(directory string, files []os.FileInfo, maxFilesInDir int) (err error) {
	var c int
	var filename string

	for _, f := range files {
		filename = path.Join(directory, f.Name())
		if c++; c > maxFilesInDir {
			err = os.Remove(filename)

			if err != nil {
				return
			}
		}
	}

	return
}

func printFilelist(directoryFiles sFiles) {
	for _, f := range directoryFiles {
		fmt.Printf("%s\n", f.Name())
	}
}
