package status

import (
	"Api-go/lib"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

)

type clientCallBack struct {
	method string `json:"method"`
	params interface{} `json:"params"`
}

type client struct {
	conn *websocket.Conn
	username string
}

var(
	all chan clientCallBack
	caller chan clientCallBack
)

var(
	chatUsers []client
	broadcastUsers []client
	chatRoomUsers []client
)

func SetOnline(c *gin.Context)  {
	ws, err := lib.UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Abort()
		return
	}
	var loginType struct{ loginType string `json:"type"`}
	_ = c.BindJSON(&loginType)
	token := c.Request.Header.Get("Authorization")
	userName,_ := lib.ParserToken(token)
	user :=  client{
		username: userName,
		conn: ws,
	}
	switch loginType.loginType {
	case "chat":
		if notIn(user,chatUsers){
			all <- clientCallBack{
				method: "GetChatUsers",
				params: []string{userName},
			}
			caller <- clientCallBack{
				method: "GetChatUsers",
				params: chatUsers,
			}
			chatUsers[len(chatUsers)] = user
		}
		break
	case "broadcast":
		if notIn(user,chatUsers){
			all <- clientCallBack{
				method: "GetBroadcastUsers",
				params: []string{userName},
			}
			caller <- clientCallBack{
				method: "GetBroadcastUsers",
				params: broadcastUsers,
			}
			chatUsers[len(broadcastUsers)] = user
		}
		break
	case "chatroom":
		if notIn(user,chatRoomUsers){
			all <- clientCallBack{
				method: "GetChatRoomUsers",
				params: []string{userName},
			}
			caller <- clientCallBack{
				method: "GetChatRoomUsers",
				params: chatRoomUsers,
			}
			chatUsers[len(chatRoomUsers)] = user
		}
		break
	}
}

func notIn(user client,users []client) bool {
	for _,cli := range users{
		if user.username == cli.username{
			return false
		}
	}
	return true
}
