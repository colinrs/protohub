package role

import (
	"net/http"

	"github.com/colinrs/protohub/pkg/httpy"

	"github.com/colinrs/protohub/internal/logic/role"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
)

func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoleListRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		l := role.NewGetRoleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleList(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
