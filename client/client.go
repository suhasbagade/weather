package client

import (
	"fmt"
	"net/url"
)

var Audience = "Super World"

func hello() {
	fmt.Println("Hello World")
}

func Getrequest(suburl string, queryparams string) {

	baseurl := url.URL{
		Scheme: "https",
		Host:   "dataservice.accuweather.com",
	}

	fmt.Println(baseurl)
	fmt.Println(suburl)
	fmt.Println(queryparams)

	//parameters
	// queryparams

	//contains
	// base url http://dataservice.accuweather.com
	// apikey AUBP88JPWuGjKkFWGiqvMcUmBTyxps39
	// GET operation
}
