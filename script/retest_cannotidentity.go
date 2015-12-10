package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ua-parser/uap-go/uaparser"
)

type uainfo struct {
	brand  string
	system string
}

func main() {
	// The "New" function will create a new UserAgent object and it will parse
	// the given string. If you need to parse more strings, you can re-use
	// this object and call: ua.Parse("another string")
	file, err := os.Open("cannot-identify") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	r.Comma = ','
	a1 := ""
	a2 := ""
	a1 = a2
	a2 = a1

	regexFile := "/Users/plutoshe/PlutoShe/Program/golang/src/github.com/ua-parser/uap-go/uap-core/regexes.yaml"
	parser, err := uaparser.New(regexFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for i := 0; ; i++ {
		log.Println(i)
		record, err := r.Read()
		// fmt.Println(i, record)
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Fatal(err)
		}
		a1 = record[0]
		client := parser.Parse(a1)
		fmt.Println("========================================")
		fmt.Println(a1)
		fmt.Println("UserAgent: " + client.UserAgent.ToString())
		fmt.Println("OS: " + client.Os.ToString())
		fmt.Println("Device: " + client.Device.ToString())
	}
}
