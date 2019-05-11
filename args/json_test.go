package args

import (
	"encoding/json"
	"fmt"
	"testing"
)

// query : .test
func TestJsonRead1(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": "hakata",
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.(string)
	if r != "hakata" {
		t.Fatal("value isn't hakata")
	}
}

// query : .test[0]
func TestJsonRead2(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": "takashi",
																"age": 19
															},
															{
																"name": "keita",
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: true,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.(map[string]interface{})

	fmt.Println(r["age"])
	a := r["age"].(float64)
	fmt.Println(a)

	if int(a) != 19 {
		t.Fatal("age isn't 19")
	}

	n, _ := r["name"].(string)
	if n != "takashi" {
		t.Fatal("name isn't takashi")
	}
}

// query : .test[0].name
func TestJsonRead3(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": "takashi",
																"age": 19
															},
															{
																"name": "keita",
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: true,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.(string)
	if r != "takashi" {
		t.Fatal("value isn't takashi")
	}
}

// query : .test[0].name.first
func TestJsonRead4(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": {
																	"first": "aaaa",
																	"second": "bbbb"
																},
																"age": 19
															},
															{
																"name": {
																	"first": "cccc",
																	"second": "dddd"
																},
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: true,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	s3 := scanResult{
		arg:     "first",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.(string)
	if r != "aaaa" {
		t.Fatal("value isn't aaaa")
	}
}

// query : .test[].name.first
func TestJsonRead5(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": {
																	"first": "aaaa",
																	"second": "bbbb"
																},
																"age": 19
															},
															{
																"name": {
																	"first": "cccc",
																	"second": "dddd"
																},
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: false,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	s3 := scanResult{
		arg:     "first",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.([]string)
	fmt.Println(r)
	m, _ := js.statement.([]interface{})
	fmt.Println(m)
	fmt.Println(m[0])
	fmt.Println(m[1])
	m0 := m[0].(string)
	if m0 != "aaaa" {
		t.Fatal("value isn't aaaa")
	}
	m1 := m[1].(string)
	if m1 != "cccc" {
		t.Fatal("value isn't cccc")
	}
	if len(m) != 2 {
		t.Fatal("count isn't 2")
	}
	// if r[0] != "aaaa" {
	// 	t.Fatal("value isn't aaaa")
	// }
	// if r[1] != "cccc" {
	// 	t.Fatal("value isn't cccc")
	// }
	// if len(r) != 2 {
	// 	t.Fatal("count isn't 2")
	// }
}

// query : .test[0].name[1].first
func TestJsonRead6(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": [
																	{
																		"first": "aaaa",
																		"second": "bbbb"
																	},
																	{
																		"first": "zzzz",
																		"second": "xxxx"	
																	}
																],
																"age": 19
															},
															{
																"name": [
																	{
																		"first": "cccc",
																		"second": "dddd"
																	},
																	{
																		"first": "yyyy",
																		"second": "eeee"	
																	}
																],
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: true,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: true,
		isIndex: true,
		index:   1,
	}
	s3 := scanResult{
		arg:     "first",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.(string)
	if r != "zzzz" {
		t.Fatal("value isn't zzzz")
	}
}

// query : .test[0].name[].second
func TestJsonRead7(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": [
																	{
																		"first": "aaaa",
																		"second": "bbbb"
																	},
																	{
																		"first": "zzzz",
																		"second": "xxxx"	
																	}
																],
																"age": 19
															},
															{
																"name": [
																	{
																		"first": "cccc",
																		"second": "dddd"
																	},
																	{
																		"first": "yyyy",
																		"second": "eeee"	
																	}
																],
																"age": 31
															}
														],
														"data": 1,
														"name": "john"
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: true,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: true,
		isIndex: false,
		index:   0,
	}
	s3 := scanResult{
		arg:     "second",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	m, _ := js.statement.([]interface{})
	fmt.Println(m)
	fmt.Println(m[0])
	fmt.Println(m[1])
	m0 := m[0].(string)
	if m0 != "bbbb" {
		t.Fatal("value isn't bbbb")
	}
	m1 := m[1].(string)
	if m1 != "xxxx" {
		t.Fatal("value isn't xxxx")
	}
	if len(m) != 2 {
		t.Fatal("count isn't 2")
	}
}

// query : .test[].name[].second
func TestJsonRead8(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": [
																	{
																		"first": "aaaa",
																		"second": "bbbb"
																	},
																	{
																		"first": "zzzz",
																		"second": "xxxx"
																	}
																],
																"age": 19
															},
															{
																"name": [
																	{
																		"first": "cccc",
																		"second": "dddd"
																	},
																	{
																		"first": "yyyy",
																		"second": "eeee"
																	}
																],
																"age": 31
															}
														],
														"data": 1
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: false,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: true,
		isIndex: false,
		index:   0,
	}
	s3 := scanResult{
		arg:     "second",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	//  success
	// r, _ := js.statement.([]interface{})
	// fmt.Println(r[0])
	// fmt.Println(r[1])
	// fmt.Println(r[0].([]interface{})[0].(map[string]interface{})["second"])

	// failure
	//fmt.Println(r[0].([]map[string]interface{})[1]["second"])

	m, _ := js.statement.([]interface{})
	_m := m[0].(string)
	if _m != "bbbb" {
		t.Fatal("value isn't bbbb")
	}
	_m = m[1].(string)
	if _m != "xxxx" {
		t.Fatal("value isn't xxxx")
	}
	_m = m[2].(string)
	if _m != "dddd" {
		t.Fatal("value isn't dddd")
	}
	_m = m[3].(string)
	if _m != "eeee" {
		t.Fatal("value isn't eeee")
	}
	if len(m) != 4 {
		t.Fatal("count isn't 4")
	}
}

// query : .test[].name[1].second
func TestJsonRead9(t *testing.T) {
	var v interface{}
	json.Unmarshal([]byte(`{
  													"test": [
															{
																"name": [
																	{
																		"first": "aaaa",
																		"second": "bbbb"
																	},
																	{
																		"first": "zzzz",
																		"second": "xxxx"	
																	}
																],
																"age": 19
															},
															{
																"name": [
																	{
																		"first": "cccc",
																		"second": "dddd"
																	},
																	{
																		"first": "yyyy",
																		"second": "eeee"	
																	}
																],
																"age": 31
															}
														],
														"data": 1
													}`), &v)
	s := scanResult{
		arg:     "test",
		isArray: true,
		isIndex: false,
		index:   0,
	}
	s2 := scanResult{
		arg:     "name",
		isArray: true,
		isIndex: true,
		index:   1,
	}
	s3 := scanResult{
		arg:     "second",
		isArray: false,
		isIndex: false,
		index:   0,
	}
	ss := []scanResult{}
	ss = append(ss, s)
	ss = append(ss, s2)
	ss = append(ss, s3)
	js := jsonState{
		statement:   v,
		queryResult: ss,
		isArray:     false,
	}
	js.jsonread()
	fmt.Println(js.statement)
	if js.statement == nil {
		t.Fatal("statement is nil")
	}
	r, _ := js.statement.([]interface{})
	m := r[0].(string)
	if m != "dddd" {
		t.Fatal("error")
	}
	m = r[1].(string)
	if m != "eeee" {
		t.Fatal("error")
	}
}
