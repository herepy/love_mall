package logic

import (
	"context"
	"love_mall/user/internal/model"
	"love_mall/utils"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddressLogic {
	return &DelAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAddressLogic) DelAddress(req *types.DelAddressRequest) error {
	err := model.DelAddress(req.UserId, req.Id)
	if err != nil {
		err = utils.NewCodeError(utils.CodeInternal)
	}

	return err
}
