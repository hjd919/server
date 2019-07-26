package main

import (
	"fmt"

	"github.com/hjd919/server/api"
	"github.com/hjd919/server/model"
)

func main() {
	user := api.UserAddReq{}
	fmt.Println(user)
	u := model.User{}
	fmt.Println(u)
}
