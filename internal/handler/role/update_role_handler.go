package role

import (
	"net/http"

	"github.com/colinrs/protohub/pkg/code"
	"github.com/colinrs/protohub/pkg/httpy"

	"github.com/colinrs/protohub/internal/logic/role"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if req.ID == 0 {
			httpx.ErrorCtx(r.Context(), w, code.ErrParam)
			return
		}
		l := role.NewUpdateRoleLogic(r.Context(), svcCtx)
		err := l.UpdateRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
