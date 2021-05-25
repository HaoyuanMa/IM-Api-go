package model

type ClientCall struct {
	Method string            `json:"method"` //客户端调用服务器的方法名
	Params map[string]string `json:"params"` //对应的参数表
}
