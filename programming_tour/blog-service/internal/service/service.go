package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
// 在应用分层中，service层主要是对业务逻辑进行封装，如果有一些业务聚合和处理可以在该层进行编码，则可以较好地隔离上下两层的逻辑