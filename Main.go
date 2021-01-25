package main
import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

// Hello returns welcome greeting
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println("Main")
}
