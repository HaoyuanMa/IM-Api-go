package socket

import (
	"Api-go/lib"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

var (
	ChatUsers      = make(map[string]*websocket.Conn)
	BroadcastUsers = make(map[string]*websocket.Conn)
	ChatRoomUsers  = make(map[string]*websocket.Conn)
	AllUsers       = make(map[string]*websocket.Conn)
)

type Message struct {
	Type        string   `json:"type"`
	From        string   `json:"from"`
	To          []string `json:"to"`
	ContentType string   `json:"contentType"`
	Content     string   `json:"content"`
}

type ClientCall struct {
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
}

type ClientCallBack struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

//var (
//	ChatChan      = make(chan ClientCallBack)
//	BroadcastChan = make(chan ClientCallBack)
//	ChatRoomChan  = make(chan ClientCallBack)
//)
//
//var (
//	Users = make(map[string]*websocket.Conn)
//)

func BuildConnection(c *gin.Context) {
	ws, err := lib.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Abort()
		return
	}
	token := c.Query("token")
	user, _ := lib.ParserToken(token)
	go Listen(ws, user)
}

func Listen(ws *websocket.Conn, user string) {
	//lastPongTime := time.Now()
	log.Printf("%s is listening\n", user)
	var file *os.File
	isUploading := false
	var ByteChan chan []byte
	StreamChan := make(chan string, 10)
	//监听处理uploadstream
	go SendStream(StreamChan)

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

		msgType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if msgType == websocket.TextMessage {
			var call ClientCall
			//token valid
			if err := json.Unmarshal([]byte(string(message)), &call); err != nil {
				err := ws.WriteMessage(msgType, []byte("Unmarshal Failed"))
				if err != nil {
					break
				}
				continue
			}

			//disconnection callback
			ws.SetCloseHandler(func(code int, text string) error {
				//remove User
				delete(AllUsers, user)
				callBack := ClientCallBack{
					Method: "RemoveUser",
					Params: user,
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
				close(StreamChan)
				log.Printf("%s close", user)
				return nil
			})
			//listen client func call
			switch call.Method {
			case "SetOnline":
				SetOnline(ws, user, call.Params)
				break
			case "SendMessage":
				go SendMessage(ws, user, call.Params)
				break
			case "StartUploadFile":
				isUploading = true
				fileName := call.Params["file"]
				userName := call.Params["user"]
				fileDir := "C:/Users/mahaoyuan/Desktop/RealTimeWeb/Api-go/UploadFiles/" + userName
				if _, err := os.Stat(fileDir); os.IsNotExist(err) {
					_ = os.Mkdir(fileDir, 0666)
				}
				fileDir += ("/" + fileName)
				if _, err := os.Stat(fileDir); !os.IsNotExist(err) {
					_ = os.Remove(fileDir)
				}
				var err error
				file, err = os.Create(fileDir)
				if err != nil {
					break
				}
				file.Close()
				file, err = os.OpenFile(fileDir, os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					break
				}
				ByteChan = make(chan []byte)
				go ReceiveFile(file, ByteChan)
				log.Printf("%s start uploading\n", user)
				break
			case "StopUploadFile":
				isUploading = false
				close(ByteChan)
				log.Printf("%s stop uploading\n", user)
				break
			case "UploadStream":
				if len(StreamChan) == cap(StreamChan) {
					_ = <-StreamChan
				}
				StreamChan <- call.Params["content"]
				break
			default:
				break
			}
		} else if msgType == websocket.BinaryMessage && isUploading {
			ByteChan <- message
		}

	}
}

func getChatUsers() []string {
	users := make([]string, 0, 1000)
	for user := range ChatUsers {
		users = append(users, user)
	}
	return users
}

func getBroadcastUsers() []string {
	users := make([]string, 0, 1000)
	for user := range BroadcastUsers {
		users = append(users, user)
	}
	return users
}

func getChatRoomUsers() []string {
	users := make([]string, 0, 1000)
	for user := range ChatRoomUsers {
		users = append(users, user)
	}
	return users
}
