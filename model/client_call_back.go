package model

type ClientCallBack struct {
	Method string      `json:"method"` //服务器调用客户端的方法名
	Params interface{} `json:"params"` //对应的参数表
}
