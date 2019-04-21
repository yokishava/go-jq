package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	path       = flag.String("p", "", "json file path")
	key        = flag.String("k", "", "json key to get value")
	outputType = flag.String("t", "string", "output type of value")
)

func main() {

	flag.Parse()
	//fmt.Println(*path, *key, *outputType)

	obj, err := readJSONFile(*path)
	if err != nil {
		errMessage := "ERROR : " + err.Error()
		fmt.Fprintln(os.Stderr, errMessage)
		os.Exit(1)
		return
	}

	value, err := getValueOfKey(obj, *key)
	if err != nil {
		errMessage := "ERROR : " + err.Error()
		fmt.Fprintln(os.Stderr, errMessage)
		os.Exit(1)
		return
	}

	fmt.Println(value)
	os.Exit(0)
}
