package handler

import (
	"love_mall/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"love_mall/user/internal/logic"
	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"
)

func PhoneLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.Response(r.Context(), w, nil, utils.NewCodeError(utils.CodeParamError))
			return
		}

		l := logic.NewPhoneLoginLogic(r.Context(), svcCtx)
		resp, err := l.PhoneLogin(&req)
		utils.Response(r.Context(), w, resp, err)
	}
}
