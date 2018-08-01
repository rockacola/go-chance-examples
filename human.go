package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("hi")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82859999))
}
