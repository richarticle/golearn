package main

//go:generate go-bindata -pkg auto -prefix views/ -o views/auto/bindata_auto.go views/templates/ views/bind/...

import (
	"fmt"
	"net/http"

	"github.com/richarticle/golearn/gobindata/views/auto"
)

func main() {
	http.Handle("/normal/", http.StripPrefix("/normal/", http.FileServer(http.Dir("views/normal"))))
	http.Handle("/bind/", http.StripPrefix("/bind/", http.FileServer(auto.AssetFS("bind"))))

	tpl := auto.LoadTemplates()
	http.HandleFunc("/template", func(w http.ResponseWriter, req *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", map[string]interface{}{"Content": "This is the content."})
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
