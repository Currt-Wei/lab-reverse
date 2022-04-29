package model

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

type ElectricMeterData struct {
	Voltage Voltage `json:"voltage"`
	Current Current `json:"current"`
	Active_power Active_power `json:"active_power"`
	Reactive_power Reactive_power `json:"reactive_power"`
	Apparent_power Apparent_power `json:"apparent_power"`
	Factor Factor `json:"factor"`
	Angel Angel `json:"angel"`
	Neutral int `json:"neutral"`
	Frequency int `json:"frequency"`
	Temperature int `json:"temperature"`
}

type ElectricMeter struct {
	VoltageA int `json:"voltageA" gorm:"column:voltageA"`
	VoltageB int `json:"voltageB" gorm:"column:voltageB"`
	VoltageC int `json:"voltageC" gorm:"column:voltageC"`
	CurrentA int `json:"currentA" gorm:"column:currentA"`
	CurrentB int `json:"currentB" gorm:"column:currentB"`
	CurrentC int `json:"currentC" gorm:"column:currentC"`
	Active_powerTotal int `json:"active_powerTotal" gorm:"column:active_powerTotal"`
	Active_powerA int `json:"active_powerA" gorm:"column:active_powerA"`
	Active_powerB int `json:"active_powerB" gorm:"column:active_powerB"`
	Active_powerC int `json:"active_powerC" gorm:"column:active_powerC"`
	Reactive_powerTotal int `json:"reactive_powerTotal" gorm:"column:reactive_powerTotal"`
	Reactive_powerA int `json:"reactive_powerA" gorm:"column:reactive_powerA"`
	Reactive_powerB int `json:"reactive_powerB" gorm:"column:reactive_powerB"`
	Reactive_powerC int `json:"reactive_powerC" gorm:"column:reactive_powerC"`
	Apparent_powerTotal int `json:"apparent_powerTotal" gorm:"column:apparent_powerTotal"`
	Apparent_powerA int `json:"apparent_powerA" gorm:"column:apparent_powerA"`
	Apparent_powerB int `json:"apparent_powerB" gorm:"column:apparent_powerB"`
	Apparent_powerC int `json:"apparent_powerC" gorm:"column:apparent_powerC"`
	FactorTotal int `json:"factorTotal" gorm:"column:factorTotal"`
	FactorA int `json:"factorA" gorm:"column:factorA"`
	FactorB int `json:"factorB" gorm:"column:factorB"`
	FactorC int `json:"factorC" gorm:"column:factorC"`
	AngelA int `json:"angelA" gorm:"column:angelA"`
	AngelB int `json:"angelB" gorm:"column:angelB"`
	AngelC int `json:"angelV" gorm:"column:angelC"`
	Neutral int `json:"neutral" gorm:"column:neutral"`
	Frequency int `json:"frequency" gorm:"column:frequency"`
	Temperature int `json:"temperature" gorm:"column:temperature"`
}

type Card struct {
	Id     int    `gorm:"column:id" json:"id"`
	Card_id     string    `gorm:"column:card_id" json:"card_id"`
	Account    string    `gorm:"column:account" json:"account"`
}

func (a Card) TableName() string {
	return "card"
}

func (e ElectricMeter) TableName() string {
	return "electricMeter"
}

func SearchCard(card_id string)(string,error){
	var card Card
	err := DB.Where("card_id = ?", card_id).Find(&card).Error
	return card.Account,err
}

func AddCard(card Card)(error){
	var card1 Card
	DB.Where("card_id = ?", card.Card_id).Find(&card1)
	if(card1.Account==""){
		err:=DB.Create(&card)
		CardId="unknown"
		return err.Error
	} else{
		return DB.Model(&card).Where("card_id",card.Card_id).Update("account",card.Account).Error
	}

}


var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var CardTime time.Time
var CardId string
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
	Temperature float32 `json:"temperature" gorm:"column:temperature"`
	Humidity float32 `json:"humidity" gorm:"column:humidity"`
}

type DoorReq struct {
	Device_id string `json:"device_id"`
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
	Timestamp string `json:"timestamp"`
	Data Door `json:"data"`
}

type Door struct {
	Secret string `json:"secret"`
	Card_id string `json:"card_id"`
	Card_label string `json:"card_label"`
}

type LightReq struct{
	Device_id string `json:"device_id"`
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
	Timestamp string `json:"timestamp"`
	Data Light `json:"data"`
}

type Light struct {
	Pwm int `json:"pwm"`
}

type DoorLog struct {
	Device_id string `json:"device_id"`
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
	Timestamp string `json:"timestamp"`
	Data Event `json:"data"`
}

type Event struct {
	Event string `json:"event"`
	Card_id string `json:"card_id"`
	Card_label string `json:"card_label"`
	Action string `json:"action"`
}

var MqttClient mqtt.Client
var InsideWeather InnerLive
func InitElecMQTT() {


	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://8.134.75.76:51883").SetClientID("lab_reserve_system")

	opts.SetUsername("SCUT2022")
	opts.SetPassword("Scut&2022-mqtt")
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

	//30AEA427B380
	if token := MqttClient.Subscribe("/smarthome/device/sensor/temperature_humidity/30AEA427B380", 0, MyESPTHCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func InitESPTHMQTT(){

	//if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}

	// 订阅主题

	if token := MqttClient.Subscribe("/smarthome/device/sensor/temperature_humidity/30AEA4254DC4", 0, MyESPTHCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func InitESPDoorMQTT(){
	if token := MqttClient.Subscribe("/smarthome/device/control/door_log/30AEA427B380", 0, MyESPDoorCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func MyESPTHCB(c mqtt.Client,msg mqtt.Message){
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

func MyESPDoorCB(c mqtt.Client,msg mqtt.Message){
	var ans DoorLog
	json.Unmarshal([]byte(msg.Payload()),&ans)
	OpenDoor(ans.Data.Card_id)
}

func OpenDoor(card_id string){
	//if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
	account, err:=SearchCard(card_id)
	if(account==""||err!=nil){
		CardId=card_id
		return
	}

	b,err:=SearchReserve(account)
	if(b==false||err!=nil){
		return
	}

	var doorReq DoorReq
	doorReq.Device_id="th_7CDFA1B52338"
	doorReq.Device_mac="7CDFA1B52338"
	doorReq.Device_type="door"
	doorReq.Timestamp="1645770729,356270"
	var door Door
	door.Card_id="08EF1234"
	door.Secret="smart_home_39381656"
	door.Card_label="user1"
	doorReq.Data=door
	req,_:=json.Marshal(doorReq)
	// 7CDFA1B52338
	//发布消息
	token := MqttClient.Publish("/smarthome/device/control/door/30AEA427B380", 0, false, string(req))
	token.Wait()

}

func LightOn(){
	var lightReq LightReq
	lightReq.Device_id=""
	lightReq.Device_type=""
	lightReq.Device_mac=""
	lightReq.Timestamp=""
	var light Light
	light.Pwm=100
	lightReq.Data=light

	req,_:=json.Marshal(lightReq)
	// 7CDFA1B52338
	//发布消息
	token := MqttClient.Publish("/smarthome/device/control/light/mac", 0, false, string(req))
	token.Wait()
}

func LightOff(){
	var lightReq LightReq
	lightReq.Device_id=""
	lightReq.Device_type=""
	lightReq.Device_mac=""
	lightReq.Timestamp=""
	var light Light
	light.Pwm=0
	lightReq.Data=light

	req,_:=json.Marshal(lightReq)
	// 7CDFA1B52338
	//发布消息
	token := MqttClient.Publish("/smarthome/device/control/light/mac", 0, false, string(req))
	token.Wait()
}

func SaveInsideWeather() error{
	return DB.Model(&InnerLive{}).Where("id",1).Update("temperature",InsideWeather.Temperature).Update("humidity",InsideWeather.Humidity).Error

}

func GetHistoryInsideWeather() (InnerLive,error){
	var weather InnerLive
	err := DB.Where("id = ?", 1).Find(&weather).Error
	return weather,err
}

func ElecDataChange(e ElectricMeterData) ElectricMeter{
	data :=ElectricMeter{}
	data.Temperature=e.Temperature
	data.Frequency=e.Frequency
	data.Neutral=e.Neutral
	data.AngelA=e.Angel.AngelA
	data.AngelB=e.Angel.AngelB
	data.AngelC=e.Angel.AngelC
	data.FactorTotal=e.Factor.FactorTotal
	data.FactorA=e.Factor.FactorA
	data.FactorB=e.Factor.FactorB
	data.FactorC=e.Factor.FactorC
	data.Apparent_powerTotal=e.Apparent_power.PowerTotal
	data.Apparent_powerA=e.Apparent_power.Apparent_powerA
	data.Apparent_powerB=e.Apparent_power.Apparent_powerB
	data.Apparent_powerC=e.Apparent_power.Apparent_powerC
	data.Reactive_powerTotal=e.Reactive_power.PowerTotal
	data.Reactive_powerA=e.Reactive_power.Reactive_powerA
	data.Reactive_powerB=e.Reactive_power.Reactive_powerB
	data.Reactive_powerC=e.Reactive_power.Reactive_powerC
	data.Active_powerTotal=e.Active_power.PowerTotal
	data.Active_powerA=e.Active_power.Active_powerA
	data.Active_powerB=e.Active_power.Active_powerB
	data.Active_powerC=e.Active_power.Active_powerC
	data.CurrentA=e.Current.CurrentA
	data.CurrentB=e.Current.CurrentB
	data.CurrentC=e.Current.CurrentC
	data.VoltageA=e.Voltage.VoltageA
	data.VoltageB=e.Voltage.VoltageB
	data.VoltageC=e.Voltage.VoltageC
	return data
}

func ElecDataChange_back(e ElectricMeter) ElectricMeterData{
	var data ElectricMeterData
	data.Temperature=e.Temperature
	data.Frequency=e.Frequency
	data.Neutral=e.Neutral
	data.Angel.AngelA=e.AngelA
	data.Angel.AngelB=e.AngelB
	data.Angel.AngelC=e.AngelC
	data.Factor.FactorTotal=e.FactorTotal
	data.Factor.FactorA=e.FactorA
	data.Factor.FactorB=e.FactorB
	data.Factor.FactorC=e.FactorC
	data.Apparent_power.PowerTotal=e.Apparent_powerTotal
	data.Apparent_power.Apparent_powerA=e.Apparent_powerA
	data.Apparent_power.Apparent_powerB=e.Apparent_powerB
	data.Apparent_power.Apparent_powerC=e.Apparent_powerC
	data.Reactive_power.PowerTotal=e.Reactive_powerTotal
	data.Reactive_power.Reactive_powerA=e.Reactive_powerA
	data.Reactive_power.Reactive_powerB=e.Reactive_powerB
	data.Reactive_power.Reactive_powerC=e.Reactive_powerC
	data.Active_power.PowerTotal=e.Active_powerTotal
	data.Active_power.Active_powerA=e.Active_powerA
	data.Active_power.Active_powerB=e.Active_powerB
	data.Active_power.Active_powerC=e.Active_powerC
	data.Current.CurrentA=e.CurrentA
	data.Current.CurrentB=e.CurrentB
	data.Current.CurrentC=e.CurrentC
	data.Voltage.VoltageA=e.VoltageA
	data.Voltage.VoltageB=e.VoltageB
	data.Voltage.VoltageC=e.VoltageC
	return data
}

func SaveElectricMeterData(e ElectricMeter) error{
	return DB.Model(&ElectricMeter{}).Where("id",1).Updates(e).Error

}

func GetHistoryElectricMeterData() (ElectricMeter,error){
	var data ElectricMeter
	err := DB.Where("id = ?", 1).Find(&data).Error
	return data,err
}