package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Test !!")
	m1 := "my name"
	//m2 := "name"

	fmt.Println(strings.ReplaceAll(m1, "m", "NOs"))
}
