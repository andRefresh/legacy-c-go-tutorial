package main

import (
	"fmt"
	"shm"
	"time"
)

func main() {
	shm := shm.OpenShm("./build/bin/my_shared_memory.shm")

	for {
		fmt.Printf("%v\n", shm)
		time.Sleep(1 * time.Second)
	}
}
