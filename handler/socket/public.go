package socket

import (
	"Api-go/lib"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type ClientCallBack struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

var (
	ChatChan      = make(chan ClientCallBack)
	BroadcastChan = make(chan ClientCallBack)
	ChatRoomChan  = make(chan ClientCallBack)
)

var (
	ChatUsers      = make(map[string]*websocket.Conn)
	BroadcastUsers = make(map[string]*websocket.Conn)
	ChatRoomUsers  = make(map[string]*websocket.Conn)
)

type ClientCall struct {
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
}

var (
	Users = make(map[string]*websocket.Conn)
)

func BuildConnection(c *gin.Context) {
	ws, err := lib.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Abort()
		return
	}
	go Listen(ws)
}

func Listen(ws *websocket.Conn) {
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if mt == websocket.TextMessage {
			var call ClientCall
			if err := json.Unmarshal([]byte(string(message)), &call); err != nil {
				err := ws.WriteMessage(mt, []byte("Unmarshal Failed"))
				if err != nil {
					break
				}
				continue
			}
			if token, ok := call.Params["token"]; !ok {
				err := ws.WriteMessage(mt, []byte("NotFound Token"))
				if err != nil {
					break
				}
				continue
			} else if user, valided := lib.ParserToken(token); !valided {
				err := ws.WriteMessage(mt, []byte("Token Error"))
				if err != nil {
					break
				}
				continue
			} else {
				log.Printf("%s is listening\n", user)
				switch call.Method {
				case "setOnline":
					SetOnline(ws, user, call.Params)
					break

				default:
					break
				}
			}
		}
	}
}
