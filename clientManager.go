package main

import (
	"bufio"
	"go-message-consumer/clients"
	"log"
	"os"
	"strings"
	"sync"
)

const CLIENT_DATA_SIZE int = 5

type ClientManager struct {
	Clients []clients.IClient
}

func (cc *ClientManager) SetClients() {
	list, _ := cc.readList("messages.csv")
	cc.Clients = make([]clients.IClient, len(list))
	for i, client := range list {
		var newCli clients.IClient
		if client[0] == "rabbit" {
			newCli = &clients.RabbitClient{}
		}
		if client[0] == "mqtt" {
			newCli = &clients.MqttClient{}
		}
		cc.Clients[i] = newCli
		newCli.SetClient(client[:])
	}
}

func (cc *ClientManager) ConsumeClients(wg *sync.WaitGroup) {
	for _, client := range cc.Clients {
		wg.Add(1)
		go client.ConsumeMessages(wg)
	}
}

func (*ClientManager) readList(file_name string) ([][CLIENT_DATA_SIZE]string, error) {
	a := [][CLIENT_DATA_SIZE]string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), ",")
		var arr [CLIENT_DATA_SIZE]string
		copy(arr[:], slice[:CLIENT_DATA_SIZE])
		a = append(a, arr)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return a, nil
}
