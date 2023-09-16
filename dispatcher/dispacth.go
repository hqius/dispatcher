package dispatcher

var dispatch = NewDispatcher()

func RegisterStates(biz string, states map[State]Process) {
	dispatch.Register(biz, states)
}

func do(ctx DptCtx, handler Handler) Handler {
	finalHandler := handler
	// global middlewares
	if len(globalMiddleware) > 0 {
		for _, middleware := range globalMiddleware {
			finalHandler = middleware(finalHandler)
		}
	}
	// biz middleware
	if middlewares, ok := bizMiddleware[ctx.Biz()]; ok {
		for _, middleware := range middlewares {
			finalHandler = middleware(finalHandler)
		}
	}
	// state middleware
	if stateMiddleware, ok := bizStateMiddleware[ctx.Biz()]; ok {
		if middlewares, ok := stateMiddleware[ctx.State()]; ok {
			for _, middleware := range middlewares {
				finalHandler = middleware(finalHandler)
			}
		}
	}
	return finalHandler
}

func Do(ctx DptCtx, req *Request) *Response {
	return do(ctx, func(ctx_ DptCtx, req_ *Request) *Response {
		return dispatch.Dispatch(ctx_, req_)
	})(ctx, req)
}
