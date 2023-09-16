package dispatcher

type Handler func(ctx DptCtx, req *Request) *Response

type Middleware func(next Handler) Handler

var globalMiddleware []Middleware
var bizMiddleware map[string][]Middleware
var bizStateMiddleware map[string]map[State][]Middleware

func AppendMiddleware(middleware Middleware) {
	globalMiddleware = append(globalMiddleware, middleware)
}

func AppendBizMiddleware(biz string, middleware Middleware) {
	if bizMiddleware == nil {
		bizMiddleware = make(map[string][]Middleware)
	}
	if _, ok := bizMiddleware[biz]; ok {
		bizMiddleware[biz] = append(bizMiddleware[biz], middleware)
	} else {
		bizMiddleware[biz] = []Middleware{middleware}
	}
}

func AppendBizStateMiddleware(biz string, state State, middleware Middleware) {
	if bizStateMiddleware == nil {
		bizStateMiddleware = map[string]map[State][]Middleware{}
	}
	if bizStateMiddleware[biz] == nil {
		bizStateMiddleware[biz] = make(map[State][]Middleware)
	}
	if bizStateMiddleware[biz][state] == nil {
		bizStateMiddleware[biz][state] = []Middleware{}
	}
	bizStateMiddleware[biz][state] = append(bizStateMiddleware[biz][state], middleware)
}
