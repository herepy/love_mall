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

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(userId int64) (resp *types.UserInfoResponse, err error) {
	user, err := model.GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = utils.NewCodeError(utils.CodeUserNotFound)
		return
	}

	if err != nil {
		logx.Errorf("get user err:%s", err.Error())
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
