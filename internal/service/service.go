package service

import (
	"github.com/hjd919/server/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"
)

// Service service.
type Service struct {
	AC  *paladin.Map
	Dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		AC:  ac,
		Dao: dao.New(),
	}
	return s
}

// Ping ping the resource.
// func (s *Service) Ping(ctx context.Context) (err error) {
// 	return s.dao.Ping(ctx)
// }

// Close close the resource.
// func (s *Service) Close() {
// 	s.dao.Close()
// }
