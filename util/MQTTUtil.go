package util

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var Vol Voltage
var Cur Current
var ActivePower Active_power
var ReactivePower Reactive_power
var ApparentPower Apparent_power
var Fac Factor
var Ang Angel
var Neutral int
var Frequency int
var Temperature int

func MyElecCB(c mqtt.Client,msg mqtt.Message){
	//fmt.Printf("MY_TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MY_MSG: %s\n", msg.Payload())
	var ans SN1
	json.Unmarshal([]byte(msg.Payload()),&ans)
	//fmt.Println("ans:",ans)
	//fmt.Println("Timestamp:",ans.Timestamp)
	//fmt.Println("Meter_sn:",ans.Meter_sn)
	//fmt.Println("Data_type:",ans.Data_type)

	if(ans.Data_type=="voltage"){
		Vol=ans.Data
	}
	if(ans.Data_type=="current"){
		var ans SN2
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Cur=ans.Data
	}
	if(ans.Data_type=="active_power"){
		var ans SN3
		json.Unmarshal([]byte(msg.Payload()),&ans)
		ActivePower=ans.Data
	}
	if(ans.Data_type=="reactive_power"){
		var ans SN4
		json.Unmarshal([]byte(msg.Payload()),&ans)
		ReactivePower=ans.Data
	}
	if(ans.Data_type=="apparent_power"){
		var ans SN5
		json.Unmarshal([]byte(msg.Payload()),&ans)
		ApparentPower=ans.Data
	}
	if(ans.Data_type=="factor"){
		var ans SN6
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Fac=ans.Data
	}
	if(ans.Data_type=="angel"){
		var ans SN7
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Ang=ans.Data
	}
	if(ans.Data_type=="neutral"){
		var ans SN8
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Neutral=ans.Data
	}
	if(ans.Data_type=="frequency"){
		var ans SN8
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Frequency=ans.Data
	}
	if(ans.Data_type=="temperature"){
		var ans SN8
		json.Unmarshal([]byte(msg.Payload()),&ans)
		Temperature=ans.Data
	}


}

type AVA struct {
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
}

type SN1 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Voltage `json:"data"`
}

type SN2 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Current `json:"data"`
}

type SN3 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Active_power `json:"data"`
}

type SN4 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Reactive_power `json:"data"`
}

type SN5 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Apparent_power `json:"data"`
}

type SN6 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Factor `json:"data"`
}

type SN7 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data Angel `json:"data"`
}

type SN8 struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data int `json:"data"`
}

type Voltage struct{
	VoltageA int `json:"voltageA"`
	VoltageB int `json:"voltageB"`
	VoltageC int `json:"voltageC"`
}

type Current struct{
	CurrentA int `json:"currentA"`
	CurrentB int `json:"currentB"`
	CurrentC int `json:"currentC"`
}

type Active_power struct{
	PowerTotal int `json:"powerTotal"`
	Active_powerA int `json:"powerA"`
	Active_powerB int `json:"powerB"`
	Active_powerC int `json:"powerC"`
}

type Reactive_power struct{
	PowerTotal int `json:"powerTotal"`
	Reactive_powerA int `json:"powerA"`
	Reactive_powerB int `json:"powerB"`
	Reactive_powerC int `json:"powerC"`
}

type Apparent_power struct{
	PowerTotal int `json:"powerTotal"`
	Apparent_powerA int `json:"powerA"`
	Apparent_powerB int `json:"powerB"`
	Apparent_powerC int `json:"powerC"`
}

type Factor struct{
	FactorTotal int `json:"factorTotal"`
	FactorA int `json:"factorA"`
	FactorB int `json:"factorB"`
	FactorC int `json:"factorC"`
}

type Angel struct{
	AngelA int `json:"angelA"`
	AngelB int `json:"angelB"`
	AngelC int `json:"angelC"`
}

type InnerLiveResp struct {
	Device_id string `json:"device_id"`
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
	Timestamp string `json:"timestamp"`
	Data InnerLive `json:"data"`
}

type InnerLive struct {
	Temperature int `json:"temperature"`
	Humidity int `json:"humidity"`
}

var MqttClient mqtt.Client
var InsideWeather InnerLive
func InitElecMQTT() {


	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://222.201.144.170:51883").SetClientID("lab_reserve_system")

	opts.SetUsername("b3351")
	opts.SetPassword("scutb3351-mqtt")
	opts.SetKeepAlive(1 * time.Hour)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Hour)

	c := mqtt.NewClient(opts)
	MqttClient=c

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题

	if token := c.Subscribe("/smarthome/dlt645/state/running/7CDFA1D66618", 0, MyElecCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//var ava AVA
	//ava.Device_mac="7CDFA1B52338"
	//ava.Device_type="dlt645"
	//a,_:=json.Marshal(ava)
	//// 7CDFA1B52338
	////发布消息
	//token := c.Publish("/smarthome/dlt645/available", 0, false, string(a))
	//token.Wait()
	//
	//token=c.Publish("/smarthome/dlt645/state/request/sn/7CDFA1B52338", 0, false, "hello")

	//time.Sleep(6 * time.Second)

	// 取消订阅
	//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	// 断开连接
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}

func InitESPMQTT(){

	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题

	if token := MqttClient.Subscribe("/smarthome/device/sensor/temperature_humidity/7CDFA1D66618", 0, MyESPCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func MyESPCB(c mqtt.Client,msg mqtt.Message){
	//fmt.Printf("MY_TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MY_MSG: %s\n", msg.Payload())
	var ans InnerLiveResp
	json.Unmarshal([]byte(msg.Payload()),&ans)
	//fmt.Println("ans:",ans)
	//fmt.Println("Timestamp:",ans.Timestamp)
	//fmt.Println("Meter_sn:",ans.Meter_sn)
	//fmt.Println("Data_type:",ans.Data_type)
	InsideWeather=ans.Data

}
