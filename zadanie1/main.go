package main

import (
	"fmt"
    "time"
	"net/http"
    "html/template"
    //"io"
    "net/url"
    "encoding/json"
)

const (
    author = "Antoni Kędra"
    port = "8080"
)

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

type MeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

type TemplateData struct {
	Name        string
    Temperature float64
    Date        string
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWriter, req *http.Request){
    data:= TemplateData{}
    if req.Method =="POST"{
        city:= req.FormValue("city")
        fmt.Println(city)
        data.Name= city
        plCity := url.QueryEscape(city)
        geoURL := fmt.Sprintf(
            "https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=pl&format=json", plCity)
        resG, err := http.Get(geoURL)
        if err != nil {
            fmt.Printf("error making http request: %s\n", err)
            return
        }
        defer resG.Body.Close()

        fmt.Printf("client: got response!\n")
        fmt.Printf("client: status code: %d\n", resG.StatusCode)
        
        var geoStruct GeoResponse
        err = json.NewDecoder(resG.Body).Decode(&geoStruct)
        if err != nil {
            fmt.Printf("error %s\n", err)
            return
        }
        var lat, lon float64
        if len(geoStruct.Results)>0 {
            lat = geoStruct.Results[0].Latitude
            lon = geoStruct.Results[0].Longitude
        }else {
			tmpl.Execute(w, data)
			return
		}

        meteoURL := fmt.Sprintf(
            "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)
        resM, err := http.Get(meteoURL)
        if err != nil {
            fmt.Printf("error making http request: %s\n", err)
            return
        }
        defer resM.Body.Close()

        fmt.Printf("client: got response!\n")
        fmt.Printf("client: status code: %d\n", resM.StatusCode)

        var meteoStruct MeteoResponse
        err = json.NewDecoder(resM.Body).Decode(&meteoStruct)
        if err != nil {
            fmt.Printf("error %s\n", err)
            return
        }
        data.Temperature = meteoStruct.CurrentWeather.Temperature
        data.Date = meteoStruct.CurrentWeather.Time
    }
    tmpl.Execute(w, data)
}


func main() {
    fmt.Println("Start:", time.Now())
	fmt.Println("Autor:", author)
	fmt.Println("Port:", port)
    
    http.HandleFunc("/",handler)
    http.ListenAndServe(":"+port, nil)
}