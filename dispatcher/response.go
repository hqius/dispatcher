package dispatcher

type Response struct {
	success bool
	code    int
	errMsg  string
	data    map[string]interface{}
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success() bool {
	return r.success
}

func (r *Response) Code() int {
	return r.code
}

func (r *Response) Msg() string {
	return r.errMsg
}

func (r *Response) OfSuccess(data map[string]interface{}) {
	r.success = true
	r.errMsg = "SUCCESS"
	r.data = data
}

func (r *Response) OfErr(code int, errMsg string) {
	r.success = false
	r.code = code
	r.errMsg = errMsg
}

func (r *Response) WithDataErr(code int, errMsg string, data map[string]interface{}) {
	r.success = false
	r.code = code
	r.errMsg = errMsg
	r.data = data
}
