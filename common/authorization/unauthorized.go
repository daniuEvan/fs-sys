/**
 * @date: 2022/5/21
 * @desc:
 */

package authorization

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UnauthorizedHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.OkJson(w, struct {
			Code int64
			Msg  string
		}{
			Code: 401,
			Msg:  "登录失效或者未登录",
		})
	}
}
