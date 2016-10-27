package filesrotator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRotate(t *testing.T) {
	testsDir := "./tests"
	os.MkdirAll(testsDir, 0755)

	for i := 1; i < 10; i++ {
		template := fmt.Sprintf("File number #%d", i)
		f, err := os.Create(path.Join(testsDir, fmt.Sprintf("File_%d.txt", i)))
		if err != nil {
			t.Fatal(err)
		}

		f.WriteString(template)
	}

	RotateByFilename(testsDir, 5)
	cnt, err := CountFiles(testsDir)
	if err != nil {
		t.Fatal(err)
	}

	if cnt != 5 {
		t.Fail()
	}
}

func CountFiles(directory string) (cnt int, err error) {

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	for _, f := range files {
		if !f.IsDir() {
			cnt++
		}
	}

	return
}
