package socket

import (
	"os"
)

func ReceiveFile(file *os.File, ByteChan chan []byte) {
	for {
		//从管道中循环读取数据
		msg, ok := <-ByteChan
		//若管道关闭，则传输结束停止读取
		if !ok {
			break
		}
		if _, err := file.Write(msg); err != nil {
			continue
		}
	}
	//关闭文件
	err := file.Close()
	if err != nil {
		return
	}
}
