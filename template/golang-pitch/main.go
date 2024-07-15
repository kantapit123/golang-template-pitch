package main

import "fmt"

func main() {

	const name = "${{ values.name }}"

	fmt.Printf("Hello world this is pitch golang template - %s", name)
}
