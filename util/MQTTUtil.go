package util

import (
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
	fmt.Printf("MY_TOPIC: %s\n", msg.Topic())
	fmt.Printf("MY_MSG: %s\n", msg.Payload())
}

type AVA struct {
	device_type string
	device_mac string
}

type SN struct {
	timestamp string
	meter_sn string
	data_type string
	data string
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
	if token := c.Subscribe("/smarthome/dlt645/available", 0, MyCB); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 7CDFA1B52338
	//发布消息
	token := c.Publish("/smarthome/dlt645/available", 0, false, "Hello World")
	token.Wait()

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

