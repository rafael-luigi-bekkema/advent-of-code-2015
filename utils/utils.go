package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var inputDir string

func init() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	inputDir = filepath.Join(filepath.Dir(filepath.Dir(f)), "input")
}

func Must[T any](p T, err error) T {
	if err != nil {
		panic(err)
	}
	return p
}

func InputReader(day int) io.ReadCloser {
	return Must(os.Open(inputPath(day)))
}

func inputPath(day int) string {
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d.txt", inputDir, day)))
}

func Input(day int) string {
	data := Must(os.ReadFile(inputPath(day)))
	return string(data)
}

func Min(nums ...int) (n int) {
	for i, num := range nums {
		if i == 0 || num < n {
			n = num
		}
	}
	return
}
