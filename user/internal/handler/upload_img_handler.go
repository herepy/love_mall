package handler

import (
	"love_mall/utils"
	"net/http"

	"love_mall/user/internal/logic"
	"love_mall/user/internal/svc"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(r)
		utils.Response(r.Context(), w, resp, err)
	}
}
