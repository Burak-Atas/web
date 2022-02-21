package utils

import "log"

func DateError(err string) {
	log.Println("hata")
}

func DataErroor(err error) {
	log.Println(err)
}

func FileError(err error) {
	log.Println("Dosya Ousturulurken hata oluştu", err)
}
