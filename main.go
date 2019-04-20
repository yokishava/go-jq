package main

import (
	"flag"
	"fmt"
)

var (
	path       = flag.String("p", "", "json file path")
	key        = flag.String("k", "", "json key to get value")
	outputType = flag.String("t", "string", "output type of value")
)

func main() {

	flag.Parse()
	fmt.Println(*path, *key, *outputType)

}
