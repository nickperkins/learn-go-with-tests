package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// Hello - Output hello in a language of your choice
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
