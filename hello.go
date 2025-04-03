package main

import "fmt"

const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	englishWorld = "World"
	spanishWorld = "Mundo"
	frenchWorld  = "Monde"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = gettingWorld(lang)
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func gettingWorld(lang string) (world string) {
	switch lang {
	case french:
		world = frenchWorld
	case spanish:
		world = spanishWorld
	default:
		world = englishWorld
	}

	return
}

func main() {
	fmt.Println(Hello("World", "english"))
}
