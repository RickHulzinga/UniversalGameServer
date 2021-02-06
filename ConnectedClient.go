package UniversalGameServer

import (
	"github.com/gorilla/websocket"
	"nextdoormediagroup.com/UniversalGameServer/EventType"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	connection *websocket.Conn
	id         string
	channel    string
}

//Gets the id of the client
func (c *Client) GetId() string {
	return c.id
}

//Gets the current channel
func (c *Client) GetChannel() string {
	return c.channel
}

//Initalizes the client and keeps listening for new messages
func (c *Client) Initalize(conn *websocket.Conn) {
	c.id = strconv.Itoa(time.Now().Nanosecond())
	c.channel = "0"
	c.connection = conn

	connections = append(connections, c)
	normCloseFunc := c.connection.CloseHandler()

	conn.SetCloseHandler(func(code int, text string) error {
		dispatchEvent(EventType.Disconnect.ToString(), "", c)

		_ = normCloseFunc(code, text)
		c.connectionDisconnect(code, text)

		return nil
	})

	dispatchEvent(EventType.Connect.ToString(), "", c)
	Log.Infoln("Client " + conn.RemoteAddr().String() + " connected with id: " + c.id)

	defer c.connection.Close()

	for {
		_, message, err := c.connection.ReadMessage()

		if err != nil {
			Log.Warning("Error reading: ", err)
			c.Disconnect()
			break
		}

		c.onMessage(string(message))

	}

}

//When a client recives a message
func (c *Client) onMessage(data string) {
	Log.Debugln(c.id + " sent a message: " + data)
	parts := strings.Fields(data)
	if len(parts) > 1 {
		dispatchEvent(parts[0], strings.ReplaceAll(data, parts[0]+" ", ""), c)
	} else {
		dispatchEvent(data, "", c)
	}

}

//When the client disconnects
func (c *Client) connectionDisconnect(code int, text string) {

	removeConnection(findConnectionIndex(c))

	Log.Infoln("Client " + c.connection.RemoteAddr().String() + " disconnected " + strconv.Itoa(code) + " " + text)
	Log.Infoln(strconv.Itoa(len(connections)) + " clients remaining")

}

//Disconnect client
func (c *Client) Disconnect() {

	e := c.connection.Close()
	if e != nil {
		Log.Warnln("Error with closing connection " + e.Error())
	}

}

//Send a message to the client
func (c *Client) SendMessage(event string, data string) {
	dat := getAlphanumericString(event) + " " + data

	err := c.connection.WriteMessage(websocket.TextMessage, []byte(dat))
	if err != nil {
		Log.Warning("Error sending message to " + c.id + " content: " + dat)
	}
}

//Switches a client to a different channel
func (c *Client) SwitchChannel(channel string) {
	c.channel = channel
	Log.Infoln("Switched " + c.id + " to channel " + c.channel)

}

//Sets a message to the channel the client is in
func (c *Client) SendMessageToChannel(event string, data string) {
	Log.Infoln("Chan " + c.channel)
	SendMessageToChannel(c.channel, event, data)
}
