package serializer

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

//BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Code: 200,
		Data: DataList{
			Items: items,
			Total: total,
		},
		Msg: "ok",
	}
}

//TokenData 带有token的Data结构
type TokenData struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
