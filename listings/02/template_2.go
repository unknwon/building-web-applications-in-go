package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse("The value is: {{.}}")
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 获取 URL 参数的值
		val := r.URL.Query().Get("val")

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, val)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
