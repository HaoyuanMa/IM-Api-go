package socket

import (
	. "Api-go/model"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func SendMessage(ws *websocket.Conn, userName string, params map[string]string) {
	msg := params["msg"]
	var message Message
	//反序列化
	err := json.Unmarshal([]byte(msg), &message)
	if err != nil {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Unmarshal Failed"))
		if err != nil {
			return
		}
	}
	//构造客户端方法调用实例
	callBack := ClientCallBack{
		Method: "ReceiveMessage",
		Params: message,
	}
	switch message.Type {
	case "chat":
		//向特定用户发送消息以调用客户端方法
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
