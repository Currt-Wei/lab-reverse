package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Live struct {
	Province string `json:"province"`
	City     string    `json:"city"`
	Adcode     string    `json:"adcode"`
	Weather     string    `json:"weather"`
	Temperature     string    `json:"temperature"`
	Humidity     string    `json:"humidity"`
	Winddirection     string    `json:"winddirection"`
	Windpower     string    `json:"windpower"`
	Reporttime	string `json:"reporttime"`

}

type RESP struct{
	Status int `json:"status"`
	Count int `json:"count"`
	Info string `json:"info"`
	Infocode string `json:"infocode"`
	Lives []Live `json:"lives"`
}
//9eea91669958b0c9eb61fd00ef014273
//bb6cfa765968b4b11454a3265440c8a6
const key = "bb6cfa765968b4b11454a3265440c8a6"
const adcode="440100"
const citycode="020"
const extensions="base"
const OUTPUT="JSON"

func GetOutsideWeather() (Live,error){
	url:=fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%v&city=%v&extensions=%v&output=%v",key,adcode,extensions,OUTPUT)
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return Live{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var r RESP
	json.Unmarshal(body,&r)
	if r.Count==0 {
		return Live{},nil
	}
	var weather Live
	weather.Province=r.Lives[0].Province
	weather.City=r.Lives[0].City
	weather.Adcode=r.Lives[0].Adcode
	weather.Weather=r.Lives[0].Weather
	weather.Temperature=r.Lives[0].Temperature
	weather.Humidity=r.Lives[0].Humidity
	weather.Winddirection=r.Lives[0].Winddirection
	weather.Windpower=r.Lives[0].Windpower
	return weather, nil
}