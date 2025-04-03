package main

import "fmt"

const englishHelloWithPrefix = "Hello, "
const spanishHelloWithPrefix = "Hola, "

func Hello(name string, lang string) string {
	if name == "" && lang == "english" {
		name = "World"
	}

	if name == "" && lang == "spanish" {
		name = "Mundo"
	}

	if lang == "spanish" {
		return spanishHelloWithPrefix + name
	}

	return englishHelloWithPrefix + name
}

func main() {
	fmt.Println(Hello("World", "english"))
}
