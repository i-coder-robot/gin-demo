package resp

type Entity struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	TotalPage int `json:"totalPage"`
	Data interface{} `json:"data"`
}
