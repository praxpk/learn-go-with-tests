package hello

import "fmt"

const (
	engGreetingPrefix     = "Hello,"
	spanishGreetingPrefix = "Hola,"
	frenchGreetingPrefix  = "Bonjour,"
	tamilGreetingPrefix   = "Vannakam,"
	kannadaGreetingPrefix = "Namaskara,"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		return spanishGreetingPrefix
	case "French":
		return frenchGreetingPrefix
	case "Tamil":
		return tamilGreetingPrefix
	case "Kannada":
		return kannadaGreetingPrefix
	default:
		return engGreetingPrefix
	}
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", greetingPrefix(language), name)
}
