package main

import (
	"fmt"
	fb "github.com/google/flatbuffers/go"
)

func main() {
	builder := fb.NewBuilder(0)
	fmt.Println("%v", builder)
}
