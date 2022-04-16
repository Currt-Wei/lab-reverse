package model

import "lab-reverse/util"

type ElectricMeterData struct {
	Voltage util.Voltage `json:"voltage"`
	Current util.Current `json:"current"`
	Active_power util.Active_power `json:"active_power"`
	Reactive_power util.Reactive_power `json:"reactive_power""`
	Apparent_power util.Apparent_power `json:"apparent_power""`
	Factor util.Factor `json:"factor"`
	Angel util.Angel `json:"angel"`
	Neutral int `json:"neutral"`
	Frequency int `json:"frequency"`
	Temperature int `json:"temperature"`
}