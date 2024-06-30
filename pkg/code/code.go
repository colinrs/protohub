package code

var (
	OKErr      = Err{Code: 0, Msg: "success"}
	UnknownErr = Err{Code: 10004, Msg: "unknown error"}
)
