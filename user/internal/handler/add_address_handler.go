package handler

import (
	"encoding/json"
	"love_mall/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"love_mall/user/internal/logic"
	"love_mall/user/internal/svc"
	"love_mall/user/internal/types"
)

func AddAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddAddressRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.Response(r.Context(), w, nil, utils.NewCodeError(utils.CodeParamError))
			return
		}

		l := logic.NewAddAddressLogic(r.Context(), svcCtx)
		req.UserId, _ = r.Context().Value("userId").(json.Number).Int64()
		err := l.AddAddress(&req)
		utils.Response(r.Context(), w, nil, err)
	}
}
