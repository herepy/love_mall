/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/12 12:13
 */

package utils

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeError struct {
	code int
	msg  string
}

func NewCodeError(code int) *CodeError {
	msg, ok := codeMsg[code]
	if !ok {
		msg = "未知错误"
	}
	return &CodeError{
		code: code,
		msg:  msg,
	}
}

func (ce *CodeError) Error() string {
	return ce.msg
}

func (ce *CodeError) Code() int {
	return ce.code
}

func Response(ctx context.Context, w http.ResponseWriter, data interface{}, err error) {
	resp := BaseResponse{}

	if err == nil {
		resp.Msg = codeMsg[CodeSuccess]
		resp.Code = CodeSuccess
		resp.Msg = "ok"
		resp.Data = data
	} else {
		switch err.(type) {
		case *CodeError:
			resp.Msg = err.(*CodeError).Error()
			resp.Code = err.(*CodeError).Code()
		default:
			resp.Msg = codeMsg[CodeInternal]
			resp.Code = CodeInternal
		}
	}

	httpx.OkJsonCtx(ctx, w, resp)
}

func UnauthorizedCallback(w http.ResponseWriter, r *http.Request, err error) {
	Response(context.Background(), w, nil, NewCodeError(CodeNoLogin))
}
