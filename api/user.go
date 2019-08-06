package api

type UserAddReq struct {
	Username string                 `json:"username"`
	Nickname string                 `json:"nickname"`
	Avatar   string                 `json:"avatar"`
	Profile  map[string]interface{} `json:"profile"`
}

type UserAddReq2 struct {
	Test string
}
