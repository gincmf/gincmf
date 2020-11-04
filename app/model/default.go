package model

type Paginate struct {
	Data     interface{} `json:"data"`
	Current  string      `json:"current" example:"1"`
	PageSize string      `json:"page_size" example:"10"`
	Total    int64       `json:"total" example:"0"`
}

type ReturnData struct {
	Code int         `json:"code" example:"1"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
