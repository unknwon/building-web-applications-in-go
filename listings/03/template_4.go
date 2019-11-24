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
		tmpl, err := template.New("test").Parse(`
{{$name1 := "alice"}}
{{$name2 := "bob"}}
{{$age1 := 18}}
{{$age2 := 23}}

{{if eq $age1 $age2}}
	年龄相同
{{else}}
	年龄不相同
{{end}}

{{if ne $name1 $name2}}
	名字不相同
{{end}}

{{if gt $age1 $age2}}
	alice 年龄比较大
{{else}}
	bob 年龄比较大
{{end}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
