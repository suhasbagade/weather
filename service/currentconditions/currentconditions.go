package currentconditions

import (
	"net/url"
	"github.com/suhasbagade/weather/service/commons/httpclient"
	"encoding/json"
)

type metric struct {
	Value    float64
	Unit     string
	UnitType int32
}

type imperial struct {
	Value    float64
	Unit     string
	UnitType int32
}

type temperature struct {
	Metric   metric
	Imperial imperial
}

type currentconditions struct {
	LocalObservationDateTime string      `json:"LocalObservationDateTime"`
	EpochTime                int64       `json:"EpochTime"`
	WeatherText              string      `json:"WeatherText"`
	WeatherIcon              int32       `json:"WeatherIcon"`
	HasPrecipitation         bool        `json:"HasPrecipitation"`
	PrecipitationType        string      `json:"PrecipitationType"`
	IsDayTime                bool        `json:"IsDayTime"`
	Temperature              temperature `json:"Temperature"`
}

func getcurrentconditionsrequesturl(locationkey string) string {
	req := httpclient.Getbaseurl()
	var path = "currentconditions/v1/" + locationkey

	req.Path = path

	params := url.Values{}
	params.Add("details", "false")

	req.RawQuery = req.RawQuery + "&" + params.Encode()

	return req.String()

}

func parsejsonresponse(r []byte) []currentconditions {
	var res []currentconditions
	json.Unmarshal(r, &res)
	return res
}


func Getcurrentconditions(locationkey string) []currentconditions {
	url := getcurrentconditionsrequesturl(locationkey)
	res := httpclient.Getrequest(url)
	temp := parsejsonresponse(res)
	return temp
}


func GetCurrentTemperatureCelsius (locationkey string) float64 {
	res := Getcurrentconditions(locationkey)
	return res[0].Temperature.Metric.Value
}

func GetCurrentTemperatureFahrenheit (locationkey string) float64 {
	res := Getcurrentconditions(locationkey)
	return res[0].Temperature.Imperial.Value
}
