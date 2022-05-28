package email

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:interface=github.com/photowey/helloioc/internal/message.Message

type Sender struct {
}

func (sender *Sender) Channel() string {
	return "email"
}

func (sender *Sender) Content() string {
	return "Hello, world!"
}

func (sender *Sender) Target() string {
	return "photowey@gmail.com"
}
