package clients

type ClientData struct {
	Name string
	Ip   string
	Port string
	Type ClientType
}

type ClientType string

const (
	Kafka  string = "kafka"
	Rabbit string = "rabbit"
	Mqtt   string = "mqtt"
)
