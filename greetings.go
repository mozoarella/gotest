package gotest

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %v, hope you enjoy your stay!", name)
	return message
}
