package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func scanf(template string, args ...interface{}) {
	fmt.Fscanf(reader, template, args...)
}

func printf(template string, args ...interface{}) {
	fmt.Fprintf(writer, template, args...)
}
