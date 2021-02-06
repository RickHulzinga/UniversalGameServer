package UniversalGameServer

var events []eventObj

type eventObj struct {
	e string
	f func(c *Client, data string)
}

func On(event string, f func(c *Client, data string)) {

	ev := new(eventObj)
	ev.e = event
	ev.f = f
	events = append(events, *ev)

}

func dispatchEvent(event string, data string, c *Client) {

	for _, ev := range events {
		if ev.e == event {

			ev.f(c, data)
		}
	}

}
