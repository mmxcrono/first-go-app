package main
import (
	"fmt"
	"strconv"
)

const englishHelloPrefix = "Hello, "

const (
	english = 1 << iota
	spanish
	french
)

// Hello returns welcome greeting
func Hello(name string, language int) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language int) string {
	switch language {
	case spanish:
		return "Hola, "
	case french:
		return "Bonjour, "
	default:
		return "Hello, "
	}
}

func main() {
	allowedLanguages := english | spanish
	language := french
	
	if language & allowedLanguages != 0 {
		fmt.Println("Language allowed: " + strconv.Itoa(language))
	}
}
