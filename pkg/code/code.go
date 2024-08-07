package code

import "net/http"

var (
	OKErr      = Err{HTTPCode: http.StatusOK, Code: 0, Msg: "success"}
	UnknownErr = Err{HTTPCode: http.StatusOK, Code: 10004, Msg: "unknown error"}
)
