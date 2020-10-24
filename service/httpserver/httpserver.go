package httpserver

import (
    "fmt"
    "net/http"
    "time"
    "html/template"
    "currentconditions"
    com "commons"
)

type pageinfo struct {
    Title string
    Description string
    Timestamp string
    Location string
    Temperature string
    Timezone string
    Latitude string
    Longitude string
}


func pagehandler(w http.ResponseWriter, req *http.Request) {
    pi := pageinfo {}
    pi.Title = "Weather Info Page"
    pi.Description = "Displays weather info for the choosen location"
    pi.Timestamp = (time.Now()).String()

    loc := req.FormValue("Location")

    if (loc != "") {

        pi.Location = loc

        res := com.Getlocationkey(loc)
        locationkey := res[0].Key
        tc := currentconditions.GetCurrentTemperatureCelsius(locationkey)
        pi.Temperature =  fmt.Sprint(tc)

        pi.Timezone = res[0].TimeZone.Code
        pi.Latitude = fmt.Sprint(res[0].GeoPosition.Latitude)
        pi.Longitude = fmt.Sprint(res[0].GeoPosition.Longitude)
    }

    templ, err := template.ParseFiles("/Users/sbagade/go_workspace/src/httpserver/pageinfo.html")
    if (err != nil) {
        panic(err)
    }

    err = templ.Execute(w, pi)
    if (err != nil) {
        panic(err)
    }

}

func Httpserverfunc() {

    http.HandleFunc("/", pagehandler)
    http.ListenAndServe(":9090", nil)
}
