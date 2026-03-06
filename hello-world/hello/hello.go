package hello

import "fmt"

const engGreetingPrefix = "Hello,"
const spanishGreetingPrefix = "Hola,"
const frenchGreetingPrefix = "Bonjour,"
const tamilGreetingPrefix = "Vannakam,"
const kannadaGreetingPrefix = "Namaskara,"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return fmt.Sprintf("%s %s!", spanishGreetingPrefix, name)
	case "French":
		return fmt.Sprintf("%s %s!", frenchGreetingPrefix, name)
	case "Tamil":
		return fmt.Sprintf("%s %s!", tamilGreetingPrefix, name)
	case "Kannada":
		return fmt.Sprintf("%s %s!", kannadaGreetingPrefix, name)
		// default:
		// 	return fmt.Sprintf("%s %s!", engGreetingPrefix, name)
	}
	return fmt.Sprintf("%s %s!", engGreetingPrefix, name)
}
