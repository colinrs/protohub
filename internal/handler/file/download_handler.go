package file

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/file"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadLogic(r.Context(), svcCtx)
		fileContentTableModel, err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileContentTableModel.FileName) // 用来指定下载下来的文件名
		w.Header().Set("Content-Transfer-Encoding", "binary")
		_, _ = w.Write([]byte(fileContentTableModel.FileContent))
		httpx.OkJsonCtx(r.Context(), w, nil)
	}
}
