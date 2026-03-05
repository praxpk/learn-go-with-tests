package hello

import "fmt"

const engGreetingPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("%s World!", engGreetingPrefix)
	}
	return fmt.Sprintf("%s %s!", engGreetingPrefix, name)
}
