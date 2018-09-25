package main

import (
	"bufio"
	"encoding/json"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Datas struct {
	Field1 string `json:"date"`
	Field2 string `json:"yslowgrade"`
	Field3 string `json:"pagespeedgrade"`
	Field4 string `json:"totalsize"`
	Field5 string `json:"totalrequests"`
	Field6 string `json:"fullyloaded"`
}

func main() {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []Datas
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fata(error)
		}
		data = append(data, Datas{
			Field1: line[0],
			Field2: line[1],
			Field3: line[3],
			Field4: line[5],
			Field5: line[6],
			Field6: line[17],
		})
	}
	dataJson, _ := json.Marshal(people)
	fmt.Println(string(data.json))
}