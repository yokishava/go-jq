package main

import (
	"testing"
)

func TestConvertMapSuccess(t *testing.T) {
	testData := map[string]interface{}{"flag": true, "sample": 1, "test": "test"}
	_, err := convertMap(testData)
	if err != nil {
		t.Fatalf("failed TestConvertMapSuccess : %#v", err)
	}
}

func TestConvertMapError(t *testing.T) {
	testData := "aaaa"
	_, err := convertMap(testData)
	if err == nil {
		t.Fatal("failed TestConvertMapError")
	}
}

func TestGetValueSuccessBool(t *testing.T) {
	testData := map[string]interface{}{"flag": true, "sample": 1, "test": "test"}
	result, err := getValue(testData, "flag")
	if err != nil {
		t.Fatalf("failed : getValue() in TestGetValueSuccessBool : %#v", err)
	}
	r, b := result.(bool)
	if b == false {
		t.Fatalf("failed : convert to bool from interface{} in TestGetValueSuccessBool : %#v", b)
	}
	if r != true {
		t.Fatalf("failed : r is different from true in TestGetValueSuccessBool : %#v", r)
	}
}

func TestGetValueSuccessString(t *testing.T) {
	testData := map[string]interface{}{"flag": true, "sample": 1, "test": "test"}
	result, err := getValue(testData, "test")
	if err != nil {
		t.Fatalf("failed : getValue() in TestGetValueSuccessString : %#v", err)
	}
	r, b := result.(string)
	if b == false {
		t.Fatalf("failed : convert to bool from interface{} in TestGetValueSuccessString : %#v", b)
	}
	if r != "test" {
		t.Fatalf("failed : r is different from true in TestGetValueSuccessString : %#v", r)
	}
}

func TestGetValueSuccessInt(t *testing.T) {
	testData := map[string]interface{}{"flag": true, "sample": 1, "test": "test"}
	result, err := getValue(testData, "sample")
	if err != nil {
		t.Fatalf("failed : getValue() in TestGetValueSuccessInt : %#v", err)
	}
	r, b := result.(int)
	if b == false {
		t.Fatalf("failed : convert to bool from interface{} in TestGetValueSuccessInt : %#v", b)
	}
	if r != 1 {
		t.Fatalf("failed : r is different from true in TestGetValueSuccessInt : %#v", r)
	}
}

func TestGetValueError(t *testing.T) {
	testData := map[string]interface{}{"flag": true, "sample": 1, "test": "test"}
	_, err := getValue(testData, "key")
	if err == nil {
		t.Fatal("failed : TestGetValueError")
	}
}

// func TestGetVaueOfKey(t *testing.T) {

// }
