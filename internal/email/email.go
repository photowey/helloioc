package email

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:interface=message.Message

type Sender struct {
}

func (sender *Sender) Channel() string {
	return "email"
}

func (sender *Sender) Content() string {
	return "Hello, world!"
}

func (sender *Sender) Target() string {
	return "13993993939"
}
