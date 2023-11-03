package main

import (
	"fmt"
)

const korea = "Korea"
const french = "French"
const englishHelloPrefix = "Hello, "
const koreaHelloPrefix = "Annyong, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Korea":
		prefix = koreaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Korea"))
}
