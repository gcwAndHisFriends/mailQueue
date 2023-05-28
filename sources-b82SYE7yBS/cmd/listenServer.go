package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
	监听其他服务器来的事件
 */

type ListenServer struct {
	localPort string
}

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}


func (listen *ListenServer)init(){
	http.HandleFunc("/", test)
	_ = http.ListenAndServe(":"+listen.localPort, nil)
}
