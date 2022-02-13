package function

import (
	"APİ/str"
	"APİ/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var URL string = "https://www.arabam.com/ikinci-el/otomobil"

func CarFound(w http.ResponseWriter) {
	for i, _ := range str.Oto {
		fmt.Fprintln(w, i)
	}
}
func Carfound(s string) bool {
	return str.Oto[s]
}

func UrlOlustur(s string) string {
	a := URL + "/" + s
	fmt.Println(s)
	fmt.Println(a)
	time.Sleep(3 * time.Second)
	return a
}

func GetData(w http.ResponseWriter, r *http.Request, s string) {
	fmt.Println("program burada 4")

	resp, err := http.Get(s)
	fmt.Println(resp)
	utils.DataCekmeError(err.Error())
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(".inner-list").Each(func(i int, s *goquery.Selection) {
		title := s.Find("li")
		fmt.Println(title)
		//fmt.Fprintf(w, string(title.Text()))
	})
}
