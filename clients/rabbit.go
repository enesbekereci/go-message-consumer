package clients

import (
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	ClientData ClientData
	Queue      amqp.Queue
	Channel    *amqp.Channel
	Conn       *amqp.Connection
}

func (c *RabbitClient) SetClient(client []string) error {
	c.ClientData.Type = ClientType(client[0])
	c.ClientData.Ip = client[1]
	c.ClientData.Port = client[2]
	c.ClientData.Name = client[3]
	fmt.Println("New Client: " + c.ClientData.Type)
	//
	var err error
	c.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	if err == nil {
		fmt.Println("Rabbit connected")
	}
	//defer conn.Close()

	c.Channel, err = c.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer c.Channel.Close()

	c.Queue, err = c.Channel.QueueDeclare(
		c.ClientData.Name, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	return nil
}

func (c *RabbitClient) ConsumeMessages(wg *sync.WaitGroup) error {
	msgs, err := c.Channel.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")
	if err == nil {
		fmt.Println("Subscribed to topic: " + c.ClientData.Name)
	}
	for d := range msgs {
		log.Printf(c.ClientData.Name)
		log.Printf("Received a message: %s", d.Body)

		d.Ack(false)
	}
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
