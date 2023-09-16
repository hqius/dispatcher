package dispatcher

import (
	"fmt"
)

type Dispatcher struct {
	meta map[string]map[State]Process
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		meta: make(map[string]map[State]Process),
	}
}

func (d *Dispatcher) Register(key string, states map[State]Process) {
	if d.meta == nil {
		panic("dispatcher is not initialized")
	}
	d.meta[key] = states
}

func (d *Dispatcher) Dispatch(ctx DptCtx, req *Request) (rsp *Response) {
	rsp = NewResponse()
	defer func() {
		if err := recover(); err != nil {
			rsp.OfErr(ProcessPanicCode, fmt.Sprintf("err: %+v", err))
		}
	}()
	if stm, ok := d.meta[ctx.Biz()]; ok {
		if process, ok := stm[ctx.State()]; ok {
			rsp = process(ctx, req)
		} else {
			rsp.OfErr(StateNotFoundCode, "StateNotFoundCode")
		}
	} else {
		rsp.OfErr(ProcessNotFoundCode, "ProcessNotFoundCode")
	}
	return rsp
}
