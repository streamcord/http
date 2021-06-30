package objects

type Request struct {
	Endpoint string
	Method   string
	Payload  interface{}
}
