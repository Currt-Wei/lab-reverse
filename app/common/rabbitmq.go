package common

import (
	"github.com/streadway/amqp"
	"log"
	"strings"
)

//Rabbitmq 初始化rabbitmq连接
type Rabbitmq struct {
	conn *amqp.Connection
	err  error
}

func NewRabbitmq() (*Rabbitmq, error) {
	rabbitmq := &Rabbitmq{
		conn: InitRabbitMQConn(),
	}
	return rabbitmq, nil
}

func (rabbitmq *Rabbitmq) CreateQueue(queueName string) error {
	ch, err := rabbitmq.conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		queueName,    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (rabbitmq *Rabbitmq) PublishQueue(queueName string, body string) error {
	ch, err := rabbitmq.conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",    // exchange
		queueName,    // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}

func (rabbitmq *Rabbitmq) ConsumeQueue(queueName string) {
	ch, _ := rabbitmq.conn.Channel()
	defer ch.Close()

	// 获取接收消息的Delivery通道
	msgs, _ := ch.Consume(
		queueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := string(d.Body)
			info := strings.Split(msg, ":")
			SendEmail(info[0], info[1])
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
