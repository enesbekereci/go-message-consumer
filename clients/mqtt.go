package clients

import (
	"fmt"
	"log"
	"strings"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	ClientData ClientData
	mqttClient mqtt.Client
}

func (c *MqttClient) SetClient(client []string) error {
	c.ClientData.Type = ClientType(client[0])
	c.ClientData.Ip = client[1]
	c.ClientData.Port = client[2]
	c.ClientData.Name = client[3]
	fmt.Println("New Client: " + c.ClientData.Type)

	//

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", c.ClientData.Ip, c.ClientData.Port))
	opts.SetClientID("go_mqtt_client" + c.ClientData.Name)
	opts.SetUsername(strings.Split(client[4], ":")[0])
	opts.SetPassword(strings.Split(client[4], ":")[1])
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	//
	c.mqttClient = mqtt.NewClient(opts)
	{
		token := c.mqttClient.Connect() //start Connection
		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	return nil
}

func (c *MqttClient) ConsumeMessages(wg *sync.WaitGroup) error {

	topic := c.ClientData.Name
	token := c.mqttClient.Subscribe(topic, 1, func(mqttClient mqtt.Client, m mqtt.Message) {
		log.Printf(c.ClientData.Name)
		log.Printf("Received a message: %s", string(m.Payload()))
	})
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
	return nil
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("MQTT Connect lost: %v\n", err)
}
