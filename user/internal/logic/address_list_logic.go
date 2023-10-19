package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"love_mall/user/internal/model"
	"love_mall/utils"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressListLogic {
	return &AddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressListLogic) AddressList(userId int64) (resp *types.AddressListResponse, err error) {
	list, err := model.AddressList(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		logx.Errorf("get address list err:%s", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	resp = &types.AddressListResponse{}
	for _, item := range list {
		resp.List = append(resp.List, types.AddressItem{
			UserId:   item.UserId,
			PostCode: item.PostCode,
			Phone:    item.Phone,
			Name:     item.Name,
			Province: item.Province,
			City:     item.City,
			District: item.District,
			Address:  item.Address,
			Default:  item.Default,
		})
	}

	return
}
