package socket

import (
	"github.com/gorilla/websocket"
	"log"
)

func SetOnline(ws *websocket.Conn, userName string, params map[string]string) {
	loginType := params["loginType"]
	switch loginType {
	case "chat":
		if _, ok := ChatUsers[userName]; !ok {
			callBack := ClientCallBack{
				Method: "GetChatUsers",
				Params: []string{userName},
			}
			for user := range ChatUsers {
				err := ChatUsers[user].WriteJSON(callBack)
				if err != nil {
					log.Printf("client.WriteJSON error: %v", err)
					ChatUsers[user].Close()
					delete(ChatUsers, user)
				}
			}

			ws.WriteJSON(ClientCallBack{
				Method: "GetChatUsers",
				Params: getChatUsers(),
			})
			ChatUsers[userName] = ws
		}
		break
	case "broadcast":
		if _, ok := BroadcastUsers[userName]; !ok {
			callBack := ClientCallBack{
				Method: "GetBroadcastUsers",
				Params: []string{userName},
			}

			for user := range BroadcastUsers {
				err := BroadcastUsers[user].WriteJSON(callBack)
				if err != nil {
					log.Printf("client.WriteJSON error: %v", err)
					BroadcastUsers[user].Close()
					delete(BroadcastUsers, user)
				}
			}
			BroadcastUsers[userName] = ws
			ws.WriteJSON(ClientCallBack{
				Method: "GetBroadcastUsers",
				Params: getBroadcastUsers(),
			})

		}
		break
	case "chatroom":
		if _, ok := ChatRoomUsers[userName]; !ok {
			callBack := ClientCallBack{
				Method: "GetChatRoomUsers",
				Params: []string{userName},
			}

			for user := range ChatRoomUsers {
				err := ChatRoomUsers[user].WriteJSON(callBack)
				if err != nil {
					log.Printf("client.WriteJSON error: %v", err)
					ChatRoomUsers[user].Close()
					delete(ChatRoomUsers, user)
				}
			}
			ChatRoomUsers[userName] = ws
			ws.WriteJSON(ClientCallBack{
				Method: "GetChatRoomUsers",
				Params: getChatRoomUsers(),
			})

		}
		break
	default:
		return
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
