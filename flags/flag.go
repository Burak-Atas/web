package flags

import (
	"APİ/function"
	"flag"
)

var AddrWebSite = flag.String("n", "", "Name web site")
var NameClass = flag.String("c", "", "tagName(If the tag name is a class name, type it with a period.(.invide) )")
var NameFile = flag.String("nf", "index.txt", "filename where you want to save the data. default(index.txt)")

func ParseFlags() {
	flag.Parse()

	*AddrWebSite = function.CreateUrl(*AddrWebSite)
	function.GetHtml(*AddrWebSite)
}
