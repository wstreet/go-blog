package serializer

type ResponseArticle struct {
	Code  int     `json:"code" example:"200"`
	Data  Article `json:"data"`
	Msg   string  `json:"msg" example:"ok"`
	Error string  `json:"error" example:""`
}
