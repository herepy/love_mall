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

type EditInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditInfoLogic {
	return &EditInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditInfoLogic) EditInfo(req *types.EditUserInfoRequest) (resp *types.UserInfoResponse, err error) {
	user, err := model.UpdateUserInfo(req.UserId, *req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = utils.NewCodeError(utils.CodeUserNotFound)
		return
	}

	if err != nil {
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	resp = &types.UserInfoResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Province: user.Province,
		City:     user.City,
	}
	return
}
