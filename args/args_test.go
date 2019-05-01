package args

import (
	"testing"
)

func TestArgStateInit(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	q := ".item[].name"
	var i interface{}

	as.init(d, q, &i)

	//fmt.Println(as.result)

	if as.arg != q {
		t.Fatal("failed : TestArgStateInit arg")
	}

	if as.data == nil {
		t.Fatal("failed : TestArgStateInit data")
	}
}

func TestArgWhileArrayFalse(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	//q := ".item[].name"
	q := ".item.name"
	var i interface{}
	as.init(d, q, &i)

	//var i interface{}

	// if as.arg != ".item[].name" {
	// 	t.Fatal("failed TestExtractValue arg")
	// }
	if as.arg != ".item.name" {
		t.Fatal("failed TestArgWhileArrayFalse arg")
	}

	//ExtractValueで失敗した理由
	/*
		- ExtarctValueでArgStateが生成される
		- このテストメソッドで新たにArgStateを生成している
		- それぞれで異なるArgStateを見ている
	*/

	//ExtractValue(d, q, &i)
	as.argWhile()

	if len(as.scan.result) == 0 {
		t.Fatal("failed TestArgWhileArrayFalse len(result)")
	}
	if as.scan.result[0].arg != "item" {
		t.Fatal("failed TestArgWhileArrayFalse as.scan.result[0].arg")
	}
	if as.scan.result[0].isArray == true {
		t.Fatal("failed TestArgWhileArrayFalse as.scan.result[0].isArray")
	}
	if len(as.scan.result) != 2 {
		t.Fatal("failed TestArgWhileArrayFalse len(result) == 2")
	}
	if as.scan.result[1].arg != "name" {
		t.Fatal("failed TestArgWhileArrayFalse as.scan.result[1].arg")
	}
	if as.scan.result[1].isArray == true {
		t.Fatal("failed TestArgWhileArrayFalse as.scan.result[1].isArray")
	}
	//fmt.Println(as.scan.result)
}

func TestArgWhileArrayTrue(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	q := ".item[].name"
	var i interface{}
	as.init(d, q, &i)

	if as.arg != ".item[].name" {
		t.Fatal("failed TestExtractValue arg")
	}

	as.argWhile()

	if len(as.scan.result) == 0 {
		t.Fatal("failed TestArgWhileArrayTrue len(result)")
	}
	if as.scan.result[0].arg != "item" {
		t.Fatal("failed TestArgWhileArrayTrue as.scan.result[0].arg")
	}
	if as.scan.result[0].isArray == false {
		t.Fatal("failed TestArgWhileArrayTrue as.scan.result[0].isArray")
	}
	if len(as.scan.result) != 2 {
		t.Fatal("failed TestArgWhileArrayTrue len(result) == 2")
	}
	if as.scan.result[1].arg != "name" {
		t.Fatal("failed TestArgWhileArrayTrue as.scan.result[1].arg")
	}
	if as.scan.result[1].isArray == true {
		t.Fatal("failed TestArgWhileArrayTrue as.scan.result[1].isArray")
	}
	//fmt.Println(as.scan.result)
}

func TestArgWhileOnlyPeriod(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	q := "."
	var i interface{}
	as.init(d, q, &i)

	as.argWhile()

	if len(as.scan.result) != 0 {
		t.Fatal("failed TestArgWhileOnlyPeriod")
	}
}

func TestArgWhileOnlyKey(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	q := ".item"
	var i interface{}
	as.init(d, q, &i)

	as.argWhile()

	if len(as.scan.result) == 0 {
		t.Fatal("failed TestArgWhileOnlyKey len(as.scan.result)")
	}
	if as.scan.result[0].arg != "item" {
		t.Fatal("failed TestArgWhileOnlyKey as.scan.result[0].arg")
	}
}

func TestArgWhileOnlyArray(t *testing.T) {
	var as ArgState
	d := map[string]string{"flag": "true", "sample": "1", "test": "morning"}
	q := ".[]"
	var i interface{}
	as.init(d, q, &i)

	as.argWhile()

	if len(as.scan.result) == 0 {
		t.Fatal("failed TestArgWhileOnlyArray len(as.scan.result)")
	}
	if as.scan.result[0].isArray != true {
		t.Fatal("failed TestArgWhileOnlyArray as.scan.result[0].isArray")
	}
	//fmt.Println(as.scan.result[0])
}
