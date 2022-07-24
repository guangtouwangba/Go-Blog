package listener

type PubSub struct {
	Data interface{}
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

type Observer interface {
	Update()
}
