package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并添加自定义模板函数
		tmpl := template.New("test").Funcs(template.FuncMap{
			"join": strings.Join,
		})

		// 解析模板内容
		_, err := tmpl.Parse(`
{{define "list"}}
    {{join . ", "}}
{{end}}
Names: {{template "list" .names}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
