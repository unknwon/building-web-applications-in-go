package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
`)

		// 根据 URL 查询参数的值创建 Inventory 实例
		inventory := &Inventory{
			SKU:  r.URL.Query().Get("sku"),
			Name: r.URL.Query().Get("name"),
		}

		// 注意：为了简化代码逻辑，这里并没有进行错误处理
		inventory.UnitPrice, _ = strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		inventory.Quantity, _ = strconv.ParseInt(r.URL.Query().Get("quantity"), 10, 64)

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, inventory)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
