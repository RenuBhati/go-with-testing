package main

// import "fmt"

// func Hello(name string) string {
// 	if name == "" {
// 		return "Hello, World"
// 	}
// 	return "Hello, " + name
// }

// // 2
// const spanish = "Spanish"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "
// const french = "French"

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	if language == spanish {
// 		return spanishHelloPrefix + name
// 	}
// 	if language == french {
// 		return frenchHelloPrefix + name
// 	}

// 	return englishHelloPrefix + name
// }

// // 3 lets use SWITCH
// const spanish = "Spanish"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "
// const french = "French"

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	prefix := englishHelloPrefix
// 	switch language {
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	case french:
// 		prefix = frenchHelloPrefix
// 	}

// 	return prefix + name
// }

// 4 REFACTORING
const (
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	french             = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

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
