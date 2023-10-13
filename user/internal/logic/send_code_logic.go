package logic

import (
	"context"
	"fmt"
	"love_mall/utils"
	"strconv"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	code := strconv.Itoa(utils.RandInt(1000, 9999))
	logx.Infof("sms code:%s", code)
	key := fmt.Sprintf("%s-%s", utils.RandStr(5), req.Phone)

	//发短信 todo

	err = l.svcCtx.Redis.Setex(fmt.Sprintf("sms|%s", key), code, 120)
	if err != nil {
		logx.Errorf("set code into redis err:", err.Error())
		err = utils.NewCodeError(utils.CodeInternal)
	}

	resp = &types.SendCodeResponse{
		Key: key,
	}

	return
}
