package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	// 创建模板对象并解析模板内容
	tmpl, err := template.ParseFiles("template_local.tmpl")
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 渲染指定模板的内容
		err = tmpl.ExecuteTemplate(w, "template_local.tmpl", map[string]interface{}{
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
