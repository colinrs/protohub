package project

import (
	"net/http"

	"github.com/colinrs/protohub/internal/models"

	"github.com/colinrs/protohub/pkg/code"

	"github.com/colinrs/protohub/internal/logic/project"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProjectRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		if err := ValidateUpdateProject(req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := project.NewUpdateProjectLogic(r.Context(), svcCtx)
		err := l.UpdateProject(&req)
		httpy.ResultCtx(r, w, nil, err)
	}
}

func ValidateUpdateProject(req types.UpdateProjectRequest) error {
	if req.Name == "" {
		return code.ErrParam
	}
	if req.Status != models.ProjectStatusBan || req.Status != models.ProjectStatusNormal {
		return code.ErrParam
	}
	return nil
}
