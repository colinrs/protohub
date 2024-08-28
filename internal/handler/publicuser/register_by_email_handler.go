package publicuser

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/publicuser"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterByEmailRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicuser.NewRegisterByEmailLogic(r.Context(), svcCtx)
		err := l.RegisterByEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
