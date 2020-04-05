package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", TaxCal)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务启动失败")
	}
}

func TaxCal(rsp http.ResponseWriter, req *http.Request) {
	io.WriteString(rsp, "欢迎使用个税计算器")
}
