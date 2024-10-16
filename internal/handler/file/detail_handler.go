package file

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/file"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDetailRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := file.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
