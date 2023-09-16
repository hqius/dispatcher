package main

import (
	"awesome/dispatcher"
	"fmt"
	"log"
	"time"
)

func main() {
	dispatcher.RegisterStates("http", map[dispatcher.State]dispatcher.Process{
		"init": func(ctx dispatcher.DptCtx, req *dispatcher.Request) *dispatcher.Response {
			data := req.Get("num").(int)
			hello := req.Get("hello").(string)
			rsp := dispatcher.NewResponse()
			rsp.OfSuccess(map[string]interface{}{
				"data": map[string]string{
					"hello": hello,
				},
				"cal": data + 1,
			})
			fmt.Printf("rsp: %+v\n", rsp)
			return rsp
		},
	})
	dispatcher.AppendMiddleware(func(next dispatcher.Handler) dispatcher.Handler {
		return func(ctx dispatcher.DptCtx, req *dispatcher.Request) *dispatcher.Response {
			log.Printf("[%s][%s][%s]: req: %+v\n", time.Now().Format("2006 15:04:05"), ctx.Biz(), ctx.State(), req)
			rsp := next(ctx, req)
			log.Printf("[%s][%s][%s]: rsp: %+v\n", time.Now().Format("2006 15:04:05"), ctx.Biz(), ctx.State(), rsp)
			return rsp
		}
	})
	rsp := dispatcher.Do(dispatcher.NewDptCtx("http", "init"), dispatcher.NewMetaRequest(map[string]interface{}{
		"num":   1,
		"hello": "world",
	}))
	fmt.Printf("rsp: %+v", rsp)
}
