package socket

import (
	. "Api-go/model"
	"fmt"
)

func SendStream(streamChan chan string) {
	for {
		//从管道中读取数据
		data, ok := <-streamChan
		if !ok {
			break
		}
		//调用所有在线客户端的DownloadStream方法以发送流式数据
		for user := range AllUsers {
			conn := AllUsers[user]
			callBack := ClientCallBack{
				Method: "DownloadStream",
				Params: data,
			}
			err := conn.WriteJSON(callBack)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("send: " + data + " to " + user)
		}
	}
}
