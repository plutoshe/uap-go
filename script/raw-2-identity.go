package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/varstr/uaparser"
)

type uainfo struct {
	brand  string
	system string
}

func process(a1, a2 string) {
	fmt.Println(a1)
	s := uaparser.Parse(a1)
	fmt.Println(s.Device)
	fmt.Println(s.DeviceType)
	fmt.Println(s.OS)
	fmt.Println(s.Browser)
	// fmt.Println(a2)
}

func main() {
	// The "New" function will create a new UserAgent object and it will parse
	// the given string. If you need to parse more strings, you can re-use
	// this object and call: ua.Parse("another string")
	file, err := os.Open("text.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	r.Comma = ','
	a1 := ""
	a2 := ""
	a1 = a2
	a2 = a1
	for i := 0; ; i++ {
		record, err := r.Read()
		// fmt.Println(i, record)
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Fatal(err)
		}
		a1 = record[7]
		fmt.Println(a1)

	}
}
