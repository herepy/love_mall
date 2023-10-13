package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"love_mall/user/internal/model"
	"love_mall/utils"
	"strings"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLoginLogic {
	return &PhoneLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneLoginLogic) PhoneLogin(req *types.PhoneLoginRequest) (resp *types.LoginResponse, err error) {
	//验证码校验
	tmp := strings.Split(req.Key, "-")
	if len(tmp) != 2 || tmp[1] != req.Phone {
		err = utils.NewCodeError(utils.CodeCheckCodeErr)
		return
	}

	code, err := l.svcCtx.Redis.Get(fmt.Sprintf("sms|%s", req.Key))
	if err != nil {
		logx.Errorf("get code from redis err:", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
		return
	}

	if code != req.Code {
		err = utils.NewCodeError(utils.CodeCheckCodeErr)
		return
	}

	//删除缓存的验证码
	l.svcCtx.Redis.Del(fmt.Sprintf("sms|%s", req.Key))

	user, err := model.GetUserByPhone(req.Phone)
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
