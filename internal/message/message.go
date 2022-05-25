package message

type Message interface {
	Channel() string
	Content() string
	Target() string
}
