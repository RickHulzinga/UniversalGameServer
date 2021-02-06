# UniversalGameServer
Simple, quick websocket based multiplayer game server.

I've created this project to rapidly develop a multiplayer solution for unity games.
Communicating via websockets all platforms should be able to communicate with this server.

### Installing
`go get github.com/NextDoorMediaGroup/UniversalGameServer`

### Example
```go
package main

import (
	"github.com/NextDoorMediaGroup/UniversalGameServer"
	"github.com/NextDoorMediaGroup/UniversalGameServer/EventType"
)

func main() {
 
 	UniversalGameServer.On("chat", func(c *UniversalGameServer.Client, data string) {
 		c.SendMessageToChannel("chat", data)
 	})
 
 	UniversalGameServer.On(EventType.Connect.ToString(), func(c *UniversalGameServer.Client, data string) {
 		c.SendMessage("chat", "Welcome on this server")
 		UniversalGameServer.BroadcastMessage("join", "Client joined "+c.GetId()+" "+c.GetChannel())
 	})
 
 	UniversalGameServer.StartServer()
 }
 
```