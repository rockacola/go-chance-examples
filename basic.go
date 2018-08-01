package main

import (
	"fmt"
	"reflect"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Basic go-chance examples ===")

	fmt.Println("Instantiate without provided seed...")
	c := chance.NewChance()

	fmt.Println("TypeOf:", reflect.TypeOf(c))
	fmt.Println("Natural():", c.Natural())
	fmt.Println("Bool():", c.Bool())
}
