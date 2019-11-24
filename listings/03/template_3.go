package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
{{if .yIsZero}}
	除数不能为 0
{{else}}
	{{.result}}
{{end}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 获取 URL 查询参数的值
		// 注意：为了简化代码逻辑，这里并没有进行错误处理
		x, _ := strconv.ParseInt(r.URL.Query().Get("x"), 10, 64)
		y, _ := strconv.ParseInt(r.URL.Query().Get("y"), 10, 64)

		// 当 y 不为 0 时进行除法运算
		yIsZero := y == 0
		result := 0.0
		if !yIsZero {
			result = float64(x) / float64(y)
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"yIsZero": yIsZero,
			"result":  result,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
