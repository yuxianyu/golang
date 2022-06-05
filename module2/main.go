package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "200")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 记录客户端 IP
		fmt.Println(r.Host)
		// 将请求头设置设置到响应头中
		header := w.Header()
		for k, v := range r.Header {
			for i := 0; i < len(v); i++ {
				header.Add(k, v[i])
			}
		}
		// 设置VERSION
		version := flag.String("VERSION", "1.17.3", "get env VERSION")
		header.Set("VERSION", *version)
	})
	http.ListenAndServe(":8088", nil)
}
