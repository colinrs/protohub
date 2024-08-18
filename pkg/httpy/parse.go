package httpy

import (
	"net/http"

	"github.com/colinrs/protohub/pkg/code"

	validator "github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var validate = validator.New()

func Parse(r *http.Request, v any) error {
	err := httpx.Parse(r, v)
	if err != nil {
		return err
	}
	err = validate.Struct(v)
	if err != nil {
		resErr := code.ErrParam.Copy()
		resErr.Msg = err.Error()
		return resErr
	}
	return nil
}
