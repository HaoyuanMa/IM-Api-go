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
	//lastPongTime := time.Now()
	for {
		//心跳检测
		//if time.Now().Sub(lastPongTime) > time.Second*10 {
		//
		//}
		//if time.Now().Sub(lastPongTime) > time.Second*5{
		//
		//}
		//ws.SetPongHandler(func(appData string) error {
		//	lastPongTime = time.Now()
		//	return nil
		//})

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
				ws.SetCloseHandler(func(code int, text string) error {
					//remove User
					callBack := ClientCallBack{
						Method: "RemoveUsers",
						Params: []string{user},
					}
					if _, ok := ChatUsers[user]; ok {
						delete(ChatUsers, user)
						for user := range ChatUsers {
							err := ChatUsers[user].WriteJSON(callBack)
							if err != nil {
								log.Printf("client.WriteJSON error: %v", err)
							}
						}
					}
					if _, ok := BroadcastUsers[user]; ok {
						delete(BroadcastUsers, user)
						for user := range BroadcastUsers {
							err := BroadcastUsers[user].WriteJSON(callBack)
							if err != nil {
								log.Printf("client.WriteJSON error: %v", err)
							}
						}
					}
					if _, ok := ChatRoomUsers[user]; ok {
						delete(ChatRoomUsers, user)
						for user := range ChatRoomUsers {
							err := ChatRoomUsers[user].WriteJSON(callBack)
							if err != nil {
								log.Printf("client.WriteJSON error: %v", err)
							}
						}
					}
					log.Printf("%s close", user)
					return nil
				})
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
