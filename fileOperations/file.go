package fileOperations

import (
	"APİ/utils"
	"encoding/csv"
	"os"
	"strings"
)

func FileOperation(str string, fl string) {
	file, err := os.Create(fl)
	utils.FileError(err)
	defer file.Close()
	if result := strings.LastIndex(".csv", file.Name()); result == 1 {
		csvWriter(str, file)
	}
}
func csvWriter(data string, flnm *os.File) {
	str := []string{data}
	writer := csv.NewWriter(flnm)
	writer.Write(str)
}
