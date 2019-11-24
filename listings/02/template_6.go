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
		tmpl, err := template.New("test").Parse(`{{/* 打印参数的值 */}}
Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 获取 URL 查询参数的值
		// 注意：为了简化代码逻辑，这里并没有进行错误处理
		sku := r.URL.Query().Get("sku")
		name := r.URL.Query().Get("name")
		unitPrice, _ := strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		quantity, _ := strconv.ParseInt(r.URL.Query().Get("quantity"), 10, 64)

		// 调用模板对象的渲染方法，并创建一个 map[string]interface{} 类型的临时变量作为根对象
		err = tmpl.Execute(w, map[string]interface{}{
			"SKU":       sku,
			"Name":      name,
			"UnitPrice": unitPrice,
			"Quantity":  quantity,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
