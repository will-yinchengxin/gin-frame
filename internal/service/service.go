package service

import (
	"context"
	"frame/global"
	"frame/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
		dao: dao.NewDao(global.DBEngine),
	}
}