package user

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/user"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func JoinProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JoinProjectRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := user.NewJoinProjectLogic(r.Context(), svcCtx)
		err := l.JoinProject(&req)
		httpy.ResultCtx(r, w, nil, err)
	}
}
