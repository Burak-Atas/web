package main

import (
	"flag"
	"os"
)

var Car = flag.String("car", "", "")
var Number = flag.Int("sn", 1, "gitmek istediğiniz site ")

func main() {
	Argument := os.Args

	flag.Parse()

	if len(Argument) == 2 {
		switch Number {
		case 1: ,

		}
	}
}

var UrlSahibinden = "https://www.sahibinden.com/"
