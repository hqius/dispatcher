package dispatcher

import (
	"fmt"
	"net/http"
)

func main() {
	RegisterStates("http", map[State]Process{
		"Start": func(ctx DptCtx, req *Request) *Response {
			return nil
		},
		"Audit": func(ctx DptCtx, req *Request) *Response {
			return nil
		},
	})
	// 定义处理请求的函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		rsp := Do(NewDptCtx("biz", "state"), &Request{})
		fmt.Fprint(w, "Hello, World!", rsp) // 在响应中写入内容
	}

	// 注册处理函数，并启动服务器监听在指定的端口
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
