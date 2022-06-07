package handler

import (
	"net/http"

	"fs-sys/service/upload/api/internal/logic"
	"fs-sys/service/upload/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileFastUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFileFastUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.FileFastUpload()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
