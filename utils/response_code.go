/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/12 14:28
 */

package utils

type ResponseCode int

const (
	CodeSuccess      = 0
	CodeInternal     = 1
	CodeParamError   = 1000
	CodeUserNotFound = 1001
	CodeNoLogin      = 1002
	CodeCheckCodeErr = 1003
)

var codeMsg = map[int]string{
	CodeSuccess:      "success",
	CodeInternal:     "内部错误",
	CodeParamError:   "参数有误",
	CodeUserNotFound: "用户不存在",
	CodeNoLogin:      "用户未登录",
	CodeCheckCodeErr: "验证码校验失败",
}
