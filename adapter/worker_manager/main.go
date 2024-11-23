package worker_manager

type BrokerMessageConsumer interface {
	Consume()
	Produce(message map[string]interface{})
}
