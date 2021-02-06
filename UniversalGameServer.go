package UniversalGameServer

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"net/http"
	"os"
)

var Log = logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.DebugLevel,
	Formatter: &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg% \n",
	},
}

var connections []*Client

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {

		Log.Warningln("Upgrade to ws failed: ", err)
		return
	}

	client := new(Client)
	client.Initalize(c)
}

func StartServer() {

	initalizeRegex()

	flag.Parse()
	Log.Infoln("Starting UGS on " + *addr)
	http.HandleFunc("/", connect)

	Log.Fatal(http.ListenAndServe(*addr, nil))
}

func GetClients() []*Client {
	return connections
}

//Not yet implemented
//func GetClientsForChannel(channel string) []Client {
//}

//Broadcasts a message to all the clients
func BroadcastMessage(event string, data string) {
	for _, c := range connections {
		c.SendMessage(event, data)
	}
}

//Broadcasts a message to a channel
func SendMessageToChannel(channel string, event string, data string) {
	for _, c := range connections {
		if c.channel == channel {
			c.SendMessage(event, data)
		}
	}
}

//Broadcasts a message to a client
func SendMessageToClient(id string, event string, data string) {
	for _, c := range connections {
		if c.id == id {
			c.SendMessage(event, data)
		}
	}
}
