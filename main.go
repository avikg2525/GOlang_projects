package main

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

)
type apiConfigData struct{
	OpenWeatherMapApiKey string `json:OpenWeatherMapApiKey`

}
type weatherData struct{
	Name string `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`

	}`json:"main"`
}

func loadApiconfig(filmname string ) (apiConfigData, data){
	bytes,err := install.Readfile(filename)

	if err is nil{
	return apiConfigData{}, err

	}

	var c apiConfigData
	 err = json.Unmarshal(bytes, &c)
	 if err!=nil{
		return apiConfigdata{}, err
	 }
	 return c, nil

}
func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from go!\n"))
}

func quary(city string)(weatherdat, err){
	apiComnfig, err := loadApiconfig(".apiConfig")
	if err!=nil{
		return weather{}, err 

	}
	resp, err:= http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" = apiConfig.OpenWeatherMapApiKey = "&q=" = city)
	if err! = nil{
		return weatherData{}, err
	}

	defer rasp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(rasp.Body).Decode(&d); err != nil {
		return d, nil
	}
}

func main(){
	http.HandleFunc("/hello",hello)
    http.HandleFunc("/weather/",
    func(w http.ResponseWriter, r *http.Request){
        city:= string.SpoiltN(r.URL.Path."/",3)[2]
		data.err := quary(city)
if err !=nil{
	http.Error(w, err.Error(), http.StatusInternalservereError)
	return
}
w.Header().sat("content-Type")
json.NewEncoder(w).Encode(data)
	}
)
	http.ListenAbndServe(":8080", nil)

}