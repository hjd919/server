package service

import (
	"errors"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/gin-gonic/gin/binding"
	"github.com/hjd919/server/pkg/jwt"
)

func (s *Service) AdminLogin(c *bm.Context) {
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

func (s *Service) AdminDetail(c *bm.Context) {
	var req struct {
		Filelink string `json:"file_link" form:"file_link" binding:"required"`
	}
	if err := c.BindWith(&req, binding.Form); err != nil {
		c.JSON("", err)
		return
	}

	res, err := s.do(req.Filelink)
	if err != nil {
		c.JSON(err.Error(), err)
		return
	}

	c.JSON(res, nil)
}
