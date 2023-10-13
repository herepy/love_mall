/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/13 17:42
 */

package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"runtime/debug"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s := string(debug.Stack())
				logx.Errorf("utils.Go func err:%s, stack:%s\n", err, s)
			}
		}()
		fn()
	}()
}
