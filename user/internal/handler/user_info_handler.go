package handler

import (
	"encoding/json"
	"love_mall/utils"
	"net/http"

	"love_mall/user/internal/logic"
	"love_mall/user/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		userId, _ := r.Context().Value("userId").(json.Number).Int64()
		resp, err := l.UserInfo(userId)
		utils.Response(r.Context(), w, resp, err)
	}
}
