package function

import (
	"APİ/fileOperations"
	"APİ/flags"

	"APİ/utils"
	"net/http"
	"strings"
)

const (
	htp = "http://"
)

func CreateUrl(str string) string {

	if err := strings.Index(htp, str); err != -1 {
		str += htp
	}

	return str
}

func GetHtml(webSite string) {
	repsonse, err := http.Get(webSite)
	utils.DataErroor(err)
	defer repsonse.Body.Close()

	Text := repsonse.Body
	fileOperations.FileOperation()

}
