package service

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func (s *Service) ListStudent(c *bm.Context) {
	// studName := c.Query("studName")
	studList := s.dao.Test()
	env := s.ac.Get("env")
	envs, _ := env.String()
	c.JSON(studList+envs, nil)
}
