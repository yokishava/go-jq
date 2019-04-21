package main

import (
	"testing"
)

func TestIsExistSuccess(t *testing.T) {
	testData := "aaaaa"
	err := isExist(testData)
	if err != nil {
		t.Fatalf("failed TestIsExistSuccess %#v", err)
	}
}

func TestIsExistError(t *testing.T) {
	testData := ""
	err := isExist(testData)
	if err == nil {
		t.Fatal("failed TestIsExistError")
	}
}

func TestIsJSONSuccess(t *testing.T) {
	testData := "./sample.json"
	err := isJSON(testData)
	if err != nil {
		t.Fatalf("failed TestIsJSONSuccess %#v", err)
	}
}

func TestIsJSONError1(t *testing.T) {
	testData := "/usr/local/bin/sample"
	err := isJSON(testData)
	if err == nil {
		t.Fatal("failed TestIsJSONError1")
	}
}

func TestIsJSONError2(t *testing.T) {
	testData := "json.txt"
	err := isJSON(testData)
	if err == nil {
		t.Fatal("failed TestIsJSONError2")
	}
}

func TestReadFileSuccess(t *testing.T) {
	testData := "./sample.json"
	_, err := readFile(testData)
	if err != nil {
		t.Fatalf("failed TestReadFileSuccess : %#v", err)
	}
	//fmt.Println(b)
}

func TestReadFileError(t *testing.T) {
	testData := "./testing.json"
	_, err := readFile(testData)
	if err == nil {
		t.Fatal("failed TestReadFileError")
	}
}
