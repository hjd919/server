package api

import "fmt"

type UserAddReq struct {
	Username string                 `json:"username"`
	Nickname string                 `json:"nickname"`
	Avatar   string                 `json:"avatar"`
	Profile  map[string]interface{} `json:"profile"`
}

func Test() {
	fmt.Println("helloTest")
}
