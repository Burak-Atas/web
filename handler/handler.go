package handler

import (
	"APİ/function"
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	urlPathElement := strings.Split(r.URL.Path, "/")
	fmt.Println("program burada 1")
	if function.Carfound(urlPathElement[1]) {
		url := function.UrlOlustur(urlPathElement[1])
		fmt.Println("program burada 2", url)

		function.GetData(w, r, url)
		fmt.Println("program burada 3 ")

	} else {
	}
}
