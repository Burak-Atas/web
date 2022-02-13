package str

import "fmt"

var Oto = map[string]bool{
	"bmw":           true,
	"audi":          true,
	"bentley":       true,
	"mercedes-benz": true,
	"tofas":         true,
	"ford":          true,
	"opel":          true,
	"toyata":        true,
	"maseratti":     true,
	"porsche":       true,
}

type Marka struct {
	Mrk  string
	next *Marka
}

type MarkaLinkedList struct {
	head *Marka
}

func (list *MarkaLinkedList) InsertFirst(otoMarka string) {
	data := &Marka{Mrk: otoMarka}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

func (list *MarkaLinkedList) InsertLast(otoMarka string) {
	data := &Marka{Mrk: otoMarka}
	if list.head == nil {
		list.head = data
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

func (list MarkaLinkedList) Listele() {
	var tmp *Marka
	m := &Marka{}
	tmp = m
	for tmp != nil {
		fmt.Println(tmp.Mrk)
		tmp = tmp.next
		fmt.Println("for döngüsünde")

	}

}

func (list *MarkaLinkedList) SearchValue(UrlMarka string) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.Mrk == "bmw" {
			return true
		}
		current = current.next
	}
	return false
}

func (list *MarkaLinkedList) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}
	return count
}
