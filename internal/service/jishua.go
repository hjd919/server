package service

import (
	"errors"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/gin-gonic/gin/binding"
	"github.com/hjd919/server/pkg/jwt"
)

func (s *Service) JishuaGetTask(c *bm.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindWith(&req, binding.JSON); err != nil {
		c.JSON("", err)
		return
	}

	// query admin
	admin, err := s.dao.AdminCheckAuth(req.Username, req.Password)
	if err != nil {
		c.JSON("", err)
		return
	}
	if admin.Username == "" {
		c.JSON("", errors.New("账号或者密码错误"))
		return
	}

	var jwtObj jwt.Jwt
	jwtObj = &jwt.AdminJwt{UserID: admin.ID}
	token, _ := jwtObj.GenerateToken()

	c.JSON(map[string]string{"token": token}, nil)
}
