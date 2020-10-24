package main

import (
	//"fmt"
	//"commons"
	//"currentconditions"
	"github.com/suhasbagade/weather/service/httpserver"
)

func main() {

//locationkey := commons.Getlocationkey("Bengaluru")
//fmt.Println(locationkey)

//currenttemperaturecelsius := currentconditions.GetCurrentTemperatureCelsius(locationkey)

//currenttemperaturefahrenheit := currentconditions.GetCurrentTemperatureFahrenheit(locationkey)

//fmt.Println(currentconditions.GetCurrentTemperatureCelsius(locationkey))
//fmt.Println(currentconditions.GetCurrentTemperatureFahrenheit(locationkey))

httpserver.Httpserverfunc()

}
