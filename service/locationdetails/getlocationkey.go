package commons

import (
	"net/url"
	"github.com/suhasbagade/weather/service/commons/httpclient"
	"encoding/json"
	"fmt"
)
/**
type region struct {
	ID            string
	LocalizedName string
	EnglishName   string
}

type country struct {
	ID            string
	LocalizedName string
	EnglishName   string
}

type administrativearea struct {
	ID            string
	LocalizedName string
	EnglishName   string
	Level         int
	LocalizedType string
	EnglishType   string
	CountryID     string
}
**/
type timezone struct {
	Code             string
	//Name             string
	//GmtOffset        float64
	//IsDaylightSaving bool
}
/**
type metric struct {
	Value    int
	Unit     string
	UnitType int
}

type imperial struct {
	Value    int
	Unit     string
	UnitType int
}

type elevation struct {
	Metric   metric
	Imperial imperial
}
**/
type geoposition struct {
	Latitude  float64
	Longitude float64
	//Elevation elevation
}
/**
type supplementaladminareas struct {
	Level         int
	LocalizedName string
	EnglishName   string
}
**/
type Getlocationkeyresponse struct {
//	Version                int                      `json:"Version"`
	Key                    string                   `json:"Key"`
/**	Type                   string                   `json:"Type"`
	Rank                   int                      `json:"Rank"`
	LocalizedName          string                   `json:"LocalizedName"`
	EnglishName            string                   `json:"EnglishName"`
	PrimaryPostalCode      string                   `json:"PrimaryPostalCode"`
	Region                 []region                 `json:"Region"`
	Country                []country                `json:"Country"`
	AdministrativeArea     []administrativearea     `json:"AdministrativeArea"`
**/	TimeZone               timezone               	`json:"TimeZone"`
	GeoPosition            geoposition            	`json:"GeoPosition"`
/**	IsAlias                bool                     `json:"IsAlias"`
	SupplementalAdminAreas []supplementaladminareas `json:"SupplementalAdminAreas"`
	DataSets               []string                 `json:"DataSets"` **/
}

func getlocationkeyrequesturl(location string) string {
	req := httpclient.Getbaseurl()
	var path = "locations/v1/cities/search"

	req.Path = path

	params := url.Values{}
	params.Add("details", "false")
	params.Add("q", location)

	req.RawQuery = req.RawQuery + "&" + params.Encode()

	return req.String()

}

func parsejsonresponse(r []byte) []Getlocationkeyresponse {
	var res []Getlocationkeyresponse
	json.Unmarshal(r, &res)
	fmt.Println(res)
	fmt.Println(res[0].Key)
	fmt.Println(res[0].TimeZone.Code)
	fmt.Println(res[0].GeoPosition.Latitude)
	fmt.Println(res[0].GeoPosition.Longitude)
	//return res[0].Key
	return res
}

func Getlocationkey(location string) []Getlocationkeyresponse {
	url := getlocationkeyrequesturl(location)
	res := httpclient.Getrequest(url)
	responsearray := parsejsonresponse(res)
	return responsearray
}

func Gettimezone(r []Getlocationkeyresponse) string {
	return r[0].TimeZone.Code
}

func GetLatitude(r []Getlocationkeyresponse) float64 {
	return r[0].GeoPosition.Latitude
}

func GetLongitude(r []Getlocationkeyresponse) float64 {
	return r[0].GeoPosition.Longitude
}
