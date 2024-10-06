package clients

import "sync"

type IClient interface {
	SetClient(client []string) error
	ConsumeMessages(wg *sync.WaitGroup) error
}
