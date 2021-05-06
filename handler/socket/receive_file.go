package socket

import (
	"os"
)

func ReceiveFile(file *os.File, ByteChan chan []byte) {
	for {
		msg, ok := <-ByteChan
		if !ok {
			break
		}
		if _, err := file.Write(msg); err != nil {
			continue
		}
	}
	file.Close()
}
