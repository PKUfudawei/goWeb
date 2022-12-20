package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	//t := template.New("web")
	//t.Parse("<div id='templateTextDiv'>Hi,{{.name}}, {{.someStr}}</div>")
	t, _ := template.ParseFiles("./templates/index.html")
	data := map[string]string{
		"name":    "zeta",
		"someStr": "This is a beginning",
	}
	t.Execute(w, data)
	/*
		r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数
		for k, v := range r.URL.Query() {
			fmt.Println("key:", k, "value:", v[0])
		}
		for k, v := range r.PostForm {
			fmt.Println("key:", k, "value:", v[0])
		}
	*/
	//fmt.Fprintf(w, "This is a beginning")
}

func main() {
	http.HandleFunc("/", myWeb)

	staticHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/js/", staticHandle)

	fmt.Println("Server is gonna open with address http://localhost:8080 in access")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server opens with error: ", err)
	}
}
