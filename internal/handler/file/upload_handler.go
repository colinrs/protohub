package file

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/file"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := file.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
