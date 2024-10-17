package role

import (
	"net/http"

	"github.com/colinrs/protohub/pkg/httpy"

	"github.com/colinrs/protohub/internal/logic/role"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
)

func GetRoleByCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoleByCodeRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := role.NewGetRoleByCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleByCode(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
