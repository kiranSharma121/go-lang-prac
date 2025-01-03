package greetings

import "strings"

func Sayhello(name string) string {
	return "Hello," + strings.TrimSpace(name)

}
