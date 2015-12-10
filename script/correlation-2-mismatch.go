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

type info struct {
	ua, os, device string
}

var Parser *uaparser.Parser

func process(testStr string) (*uaparser.Client, info) {
	client := Parser.Parse(testStr)
	// fmt.Println(testStr)
	// fmt.Println("UserAgent: " + client.UserAgent.ToString())
	// fmt.Println("OS: " + client.Os.ToString())
	// fmt.Println("Device: " + client.Device.ToString())
	// fmt.Println("Device: " + client.Os.ToVersionString())
	return client, info{client.UserAgent.ToString(), client.Os.ToString(), client.Device.ToString()}
	// fmt.Println(a2)
}

func main() {
	regexFile := "../uap-core/regexes.yaml"
	if len(os.Args) > 1 {
		regexFile = os.Args[1]
	}
	var err error
	Parser, err = uaparser.New(regexFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// The "New" function will create a new UserAgent object and it will parse
	// the given string. If you need to parse more strings, you can re-use
	// this object and call: ua.Parse("another string")
	file, err := os.Open("ans.csv") // For read access.
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
		log.Println(i)
		record, err := r.Read()
		// fmt.Println(i, record)
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Fatal(err)
		}
		if i%2 == 0 {
			a1 = record[5]
			// fmt.Println(record)
		}
		if i%2 == 1 {
			a2 = record[5]
			w1, u1 := process(a1)
			w2, u2 := process(a2)
			if w1.Os.Family != w2.Os.Family {
				// if w1.Device.ToString() != w2.Device.ToString() {
				// 	fmt.Println(u1, u2)
				// }
			} else {
				// w1.Os.ToVersionString() != w2.Os.ToVersionString() ||
				if w1.Device.ToString() == "IOS-DEVICE" && w2.Device.ToString() == "IPHONE" {
					continue
				}
				if w1.Os.ToString() != w2.Os.ToString() && w1.Device.ToString() != w2.Device.ToString() {
					fmt.Println(a1, "\n", a2, "\n", u1.device, "\n", u2.device, "\n==============\n")
				}
			}
		}

	}
}
