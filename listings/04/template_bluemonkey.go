package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	p := bluemonday.UGCPolicy()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并添加自定义模板函数
		tmpl := template.New("test").Funcs(template.FuncMap{
			"sanitize": func(s string) template.HTML {
				return template.HTML(p.Sanitize(s))
			},
		})

		// 解析模板内容
		_, err := tmpl.Parse(`
<html>
<body>
	<p>{{.content | sanitize}}</p>
</boyd>
</html>
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"content": `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
