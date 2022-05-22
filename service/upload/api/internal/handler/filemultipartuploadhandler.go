package handler

import (
	"net/http"

	"fs-sys/service/upload/api/internal/logic"
	"fs-sys/service/upload/api/internal/svc"
	"fs-sys/service/upload/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileMultipartUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MultipartUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileMultipartUploadLogic(r.Context(), svcCtx)
		err := l.FileMultipartUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
