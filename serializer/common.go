package serializer

// Response 响应的 json 结构
// 基础序列化器
type Response struct {
	// 响应状态码
	Status int `json:"status"`
	// 返回数据
	Data interface{} `json:"data"`
	// 返回的消息
	Msg string `json:"msg"`
	// 返回的错误
	Error string `json:"error"`
}

// Token 数据的 json 结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
