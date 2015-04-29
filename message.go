package json_bench

// number of messages we want to test
const Messages = 1e5

// test embedding
type Status int32

//go:generate codecgen -o message_codec.go message.go
type Message struct {
	Body string
	Status
}
