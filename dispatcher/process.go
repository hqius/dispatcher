package dispatcher

type Process func(ctx DptCtx, req *Request) *Response
