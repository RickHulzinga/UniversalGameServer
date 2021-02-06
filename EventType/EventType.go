package EventType

type EventType string

const (
	Disconnect EventType = "disconnect"
	Connect    EventType = "connect"
)

func (e EventType) ToString() string {

	return string(e)

}
