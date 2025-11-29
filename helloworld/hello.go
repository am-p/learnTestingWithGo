package main

import (
	"fmt"
)

const (
 deutsch = "Deutsch"
 francois = "Francois"
 englishHelloPrefix = "Hello, "
 deutschHelloPrefix = "Hallo, "
 francoisBonJourPrefix = "Bon jour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case deutsch:
		prefix = deutschHelloPrefix
	case francois:
		prefix = francoisBonJourPrefix
	default:
		prefix = englishHelloPrefix
	}
	
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Ariel", "Deutsch"))
}
