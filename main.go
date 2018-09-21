package main

/*
needs:
fmt
encoding/csv
net/http
*/
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func find_in_csv(code string) string {
	var o string
	csvFile, _ := os.Open("zcta_crosswalk.csv")
	r := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if line[0] == code {
			o = line[1]
			break
		}
	}
	return o

}

func main() {
	fmt.Println(find_in_csv("01701"))
}
