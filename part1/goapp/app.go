package main

import (
	"fmt"
	"mmf"
	"time"
)

func main() {
	mmf := mmf.OpenMmf("./build/bin/memory_mapped_file.mmf")

	for {
		fmt.Printf("%v\n", mmf)
		time.Sleep(1 * time.Second)
	}
}
