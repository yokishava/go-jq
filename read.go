package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// read file

// check file path
func isExist(v string) error {
	if v == "" {
		err := &Error{
			Message: "the file path isn't written on this commands. ",
		}
		return err
	}
	return nil
}

//check file is .json
func isJSON(p string) error {
	if strings.Contains(p, ".json") == false {
		m := "cannot find json file. : " + p
		err := &Error{
			Message: m,
		}
		return err
	}
	return nil
}

//check can read json file
func readFile(p string) ([]byte, error) {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		newerr := &Error{
			Message: "no such file",
		}
		return nil, newerr
	}
	return data, err
}

//readJOSN : convert to interface{} object from json file
func readJSONFile(p string) (interface{}, error) {
	err := isExist(p)
	if err != nil {
		return nil, err
	}
	err = isJSON(p)
	if err != nil {
		return nil, err
	}
	b, e := readFile(p)
	if e != nil {
		return nil, e
	}
	var i interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// func readJson() {
// 	// data => []byte
// 	data, err := ioutil.ReadFile("./sample.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	var j interface{}
// 	//Unmarshalは構造体のポインタだけではなく、*interface{}を渡すことができる
// 	json.Unmarshal(data, &j)

// 	//fmt.Println(j)
// 	fmt.Println(j.(map[string]interface{})["go.buildOnSave"].(string))
// }
