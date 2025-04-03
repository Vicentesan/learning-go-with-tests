package main

import "fmt"

const englishHelloWithPrefix = "Hello, "
const spanishHelloWithPrefix = "Hola, "
const frenchHelloWithPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, lang string) string {
	if name == "" && lang == "english" {
		name = "World"
	}

	if name == "" && lang == spanish {
		name = "Mundo"
	}

	if name == "" && lang == spanish {
		name = "Monde"
	}

	if lang == spanish {
		return spanishHelloWithPrefix + name
	}

	if lang == french {
		return frenchHelloWithPrefix + name
	}

	return englishHelloWithPrefix + name
}

func main() {
	fmt.Println(Hello("World", "english"))
}
