package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("hi chance")

	c := chance.NewChance()
	fmt.Println("foo:", c.Natural())
}
