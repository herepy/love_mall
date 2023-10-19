package logic

import (
	"context"
	"love_mall/user/internal/model"
	"love_mall/utils"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressLogic) AddAddress(req *types.AddAddressRequest) error {
	address := model.UserAddress{
		UserId:   req.UserId,
		PostCode: req.PostCode,
		Phone:    req.Phone,
		Name:     req.Name,
		Province: req.Province,
		City:     req.City,
		District: req.District,
		Default:  req.Default,
		Status:   model.StatusOk,
	}

	err := model.AddAddress(address)
	if err != nil {
		err = utils.NewCodeError(utils.CodeInternal)
	}

	return err
}
