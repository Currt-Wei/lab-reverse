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

func MyCB(c mqtt.Client,msg mqtt.Message){
	//fmt.Printf("MY_TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MY_MSG: %s\n", msg.Payload())
	var ans SN
	json.Unmarshal([]byte(msg.Payload()),&ans)
	fmt.Println(ans.Timestamp)

}

func MyCB2(c mqtt.Client,msg mqtt.Message){
	//fmt.Printf("MY_TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MY_MSG: %s\n", msg.Payload())
	fmt.Println("hello world")

}

type AVA struct {
	Device_type string `json:"device_type"`
	Device_mac string `json:"device_mac"`
}

type SN struct {
	Timestamp string `json:"timestamp"`
	Meter_sn string	`json:"meter_sn"`
	Data_type string `json:"data_type"`
	Data string `json:"data"`
}

func InitMQTT() {


	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://222.201.144.170:51883").SetClientID("wjh_client1")

	opts.SetUsername("b3351")
	opts.SetPassword("scutb3351-mqtt")
	opts.SetKeepAlive(1 * time.Hour)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Hour)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("/smarthome/dlt645/state/info/7CDFA1B52338", 0, MyCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := c.Subscribe("/smarthome/dlt645/state/running/7CDFA1B52338", 0, MyCB2); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	var ava AVA
	ava.Device_mac="7CDFA1B52338"
	ava.Device_type="dlt645"
	a,_:=json.Marshal(ava)
	// 7CDFA1B52338
	//发布消息
	token := c.Publish("/smarthome/dlt645/available", 0, false, string(a))
	token.Wait()

	token=c.Publish("/smarthome/dlt645/state/request/sn/7CDFA1B52338", 0, false, "hello")

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

