package client

import "net/url"

const apihost = "dataservice.accuweather.com"
const apikey = "AUBP88JPWuGjKkFWGiqvMcUmBTyxps39"
const apiurlscheme = "https"

/**
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
**/

func getbaseurl() {
	baseurl := url.URL{
		Scheme: apiurlscheme,
		Host:   apihost,
	}
	return baseurl
}

func GetlocationKey() {

}

func GetRequest(path string) {

	// check for all possible paths

}
