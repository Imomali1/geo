package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var reader = bufio.NewReader(os.Stdin)

func scanf(template string, args ...interface{}) (err error) {
	_, err = fmt.Fscanf(reader, template, args...)
	return
}

func info() {
	fmt.Printf("App is running on %s/%s\n\n", runtime.GOOS, runtime.GOARCH)
}
