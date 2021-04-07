package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	size, err := DiskAvail(Fsroot())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("avail size(GB)", size/1024/1024/1024)
}
