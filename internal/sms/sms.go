package sms

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Sender struct {
}

func (sender *Sender) Channel() string {
	return "sms"
}

func (sender *Sender) Content() string {
	return "Hello, world!"
}

func (sender *Sender) Target() string {
	return "13993993939"
}
