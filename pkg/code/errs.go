package code

type Err struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Errors []*Error `json:"errors,omitempty"`
}

func (e *Err) Error() string {
	return e.Msg
}

func (e *Err) GetCode() int {
	return e.Code
}

func (e *Err) GetMsg() string {
	return e.Msg
}

func (e *Err) WithErrors(items []*Error) {
	e.Errors = append(e.Errors, items...)
}

func (e *Err) GetErrors() []*Error {
	return e.Errors
}

type Error struct {
	Attr   string `json:"attr,omitempty"`
	Code   int    `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (e *Error) Error() string {
	return e.Detail
}

func (e *Error) GetCode() int {
	return e.Code
}
