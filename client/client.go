package client

import (
	"fmt"
	"net/url"
)

func getrequest(suburl string, queryparams string) {

	baseurl := url.URL{
		Scheme: "https",
		Host:   "dataservice.accuweather.com",
	}

	fmt.Println(baseurl)

	//parameters
	// queryparams

	//contains
	// base url http://dataservice.accuweather.com
	// apikey AUBP88JPWuGjKkFWGiqvMcUmBTyxps39
	// GET operation
}
