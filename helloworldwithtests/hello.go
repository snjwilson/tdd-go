package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const exclamationMark = "!"

func hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefixer(language) + name + exclamationMark
}

func greetingPrefixer(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(hello("Sanjay","english"))
}
