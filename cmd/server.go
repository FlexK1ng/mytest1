package main

import (
	"fmt"

	"github.com/devomot/jenitter/most"
)

func init() {
	fmt.Println("Hello from init!")

}

func main() {
	foo()

}

func foo() {
	fmt.Println("Foo!")
	most.Mcad()
}
