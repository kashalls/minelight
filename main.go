package main

import "fmt"

const banner = `
sunlight server - a network tester 
version: %s (%s)

`

var (
	Version = "local"
	Gitsha  = "?"
)

func main() {
	fmt.Printf(banner, Version, Gitsha)

}
