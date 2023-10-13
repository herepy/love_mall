package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"love_mall/user/internal/model"
	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"
	"love_mall/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxLoginLogic) WxLogin(req *types.WxLoginRequest) (resp *types.LoginResponse, err error) {
	//获取openid
	openid, err := svc.GetOpenid(req.Code, l.svcCtx.Config.Wechat.Appid, l.svcCtx.Config.Wechat.Secret)
	if err != nil {
		logx.Errorf("get openid err:%s", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	user, err := model.GetUserByOpenid(openid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = utils.NewCodeError(utils.CodeUserNotFound)
		return
	}

	if err != nil {
		logx.Errorf("get user err:%s", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	token, err := utils.MakeJwtToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, user.Id)
	if err != nil {
		logx.Errorf("make jwt err:%s", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	resp = &types.LoginResponse{
		Token: token,
		UserInfo: types.UserInfoResponse{
			Id:       user.Id,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Gender:   user.Gender,
			Phone:    user.Phone,
			Province: user.Province,
			City:     user.City,
		},
	}

	return
}
