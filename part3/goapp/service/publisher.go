package service

type Publisher struct {
	subscribers map[any]chan any
}

func NewPublisher() (p Publisher) {
	p = Publisher{}
	p.subscribers = map[any]chan any{}

	return
}

func (p *Publisher) Subscribe(subscriber any) <-chan any {
	sc := make(chan any)
	p.subscribers[subscriber] = sc

	return sc
}

func (p *Publisher) Unsubscribe(subscribed any) {
	delete(p.subscribers, subscribed)
}
