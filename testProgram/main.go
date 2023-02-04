package main

import (
	"flag"
	"fmt"
)

const DEFAULT_MESSAGE = "default message"

var message string

func init() {
	flag.StringVar(&message, "message", "", "Message to print, if not using default message")
}

func main() {
	flag.Parse()

	if message != "" {
		fmt.Println(message)
	} else {
		fmt.Println(DEFAULT_MESSAGE)
	}
}
