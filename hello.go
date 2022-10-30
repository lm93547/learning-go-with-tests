package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const indo = "Indonesian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const indoHelloPrefix = "Halo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case indo:
		prefix = indoHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
