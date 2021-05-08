package socket

import "fmt"

func SendStream(streamChan chan string) {
	for {
		msg, ok := <-streamChan
		if !ok {
			break
		}
		for user := range AllUsers {
			conn := AllUsers[user]
			callBack := ClientCallBack{
				Method: "DownloadStream",
				Params: msg,
			}
			err := conn.WriteJSON(callBack)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("send: " + msg + " to " + user)
		}
	}
}
