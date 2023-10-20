package logic

import (
	"context"
	"love_mall/utils"
	"net/http"

	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(r *http.Request) (resp *types.UploadImgResponse, err error) {
	imgUrl, err := svc.UploadImg(r, l.svcCtx.Config.Oss.Key, l.svcCtx.Config.Oss.Secret, l.svcCtx.Config.Oss.Endpoint, l.svcCtx.Config.Oss.Bucket, "img")
	if err != nil {
		err = utils.NewCodeError(utils.CodeParamError)
		return
	}

	resp = &types.UploadImgResponse{}
	resp.Url = imgUrl
	return
}
