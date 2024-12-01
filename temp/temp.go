package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"price": "costly"}
	// m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
	a, _ := fmt.Printf("my name is %s\n", "renu")
	fmt.Println(a)

}
