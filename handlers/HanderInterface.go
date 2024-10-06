package handlers

type HandlerInterface[T any] interface {
	Set(func(*T))
	Parse(message []byte, input *T) error
	Handle(*T)
}
