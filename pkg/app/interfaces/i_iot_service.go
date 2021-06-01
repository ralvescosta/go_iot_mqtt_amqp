package interfaces

type IIotCliente interface {
	Connect() error
	Sub(topic string, qos int, Handler func(payload []byte, topic string))
	Pub(topic string, qos int, payload []byte)
	Defer()
}
