package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"../uaparser" // You could change this to a github repo as well
)

func main() {
	flag.Parse()
	args := flag.Args()
	var testStr string
	if len(args) > 0 {
		testStr = args[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		testStr, _ = reader.ReadString('\n')
		// testStr := "Dalvik/1.6.0 (Linux; U; Android 4.4.4; Coolpad 8675-FHD MIUI/5.9.28)"
	}

	regexFile := "../uap-core/regexes.yaml"
	// if len(os.Args) > 2 {
	// 	regexFile = os.Args[2]
	// }
	parser, err := uaparser.New(regexFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	client := parser.Parse(testStr)
	fmt.Println(testStr)
	fmt.Println("UserAgent: " + client.UserAgent.ToString())
	fmt.Println("OS: " + client.Os.ToString())

	fmt.Println("Device: " + client.Device.ToString())
}
