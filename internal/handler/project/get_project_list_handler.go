package project

import (
	"net/http"

	"github.com/colinrs/protohub/internal/logic/project"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProjectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProjectListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := project.NewGetProjectListLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectList(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
