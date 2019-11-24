package main

import (
	"fmt"
	"log"
	"net/http"
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
SKU: {{.Inventory.SKU}}
Name: {{.Inventory.Name}}
UnitPrice: {{.Inventory.UnitPrice}}
Quantity: {{.Inventory.Quantity}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"Inventory": Inventory{
				SKU:       "11000",
				Name:      "Phone",
				UnitPrice: 699.99,
				Quantity:  666,
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
