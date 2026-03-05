package hello

import "fmt"

const engGreetingPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s!", engGreetingPrefix, name)
}
