package main

import (
	"fmt"

	"github.com/suhasbagade/weather/client"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(client.Audience)
	client.Hello()

	client.Getrequest("ayz", "abc")
}
