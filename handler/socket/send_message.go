package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

func SendMessage(ws *websocket.Conn, userName string, params map[string]string) {
	msg := params["msg"]
	var message Message
	if err := json.Unmarshal([]byte(string(msg)), &message); err != nil {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Unmarshal Failed"))
		if err != nil {
			return
		}
	}
	callBack := ClientCallBack{
		Method: "ReceiveMessage",
		Params: message,
	}
	switch message.Type {
	case "chat":
		to := ChatUsers[message.To[0]]
		err := to.WriteJSON(callBack)
		if err != nil {
			return
		}
		break
	case "broadcast":
		for user := range BroadcastUsers {
			err := BroadcastUsers[user].WriteJSON(callBack)
			if err != nil {
				return
			}
		}
		break
	case "chatroom":
		for user := range ChatRoomUsers {
			err := ChatRoomUsers[user].WriteJSON(callBack)
			if err != nil {
				return
			}
		}
		break
	default:
		break
	}
}
