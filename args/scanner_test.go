package args

import (
	"testing"
)

func TestInitScanner(t *testing.T) {
	var s scanner
	s.initScanner()

	if s.step == nil {
		t.Fatal("failed initScanner (step)")
	}
	if s.isArray == true {
		t.Fatal("failed initScanner (isArray)")
	}
	if s.isIndex == true {
		t.Fatal("failed initScanner (isIndex)")
	}
	if len(s.indexBytes) != 0 {
		t.Fatal("failed initScanner (indexBytes)")
	}
	if len(s.argBytes) != 0 {
		t.Fatal("failed initScanner (argBytes)")
	}
	if len(s.result) != 0 {
		t.Fatal("failed initScanner (result)")
	}
}

func TestStateBeginPeriod(t *testing.T) {
	var s scanner
	s.initScanner()

	var b byte = '.'

	stateBeginPeriod(&s, b)

	if len(s.argBytes) != 0 {
		t.Fatal("failed stateBeginPeriod")
	}
}

func TestStateBeginStringOrArray(t *testing.T) {
	var s scanner
	s.initScanner()

	var b byte = '['

	stateBeginStringOrArray(&s, b)

	if len(s.argBytes) != 0 {
		t.Fatal("failed stateBeginStringOrArray [")
	}

	b = 'i'
	stateBeginStringOrArray(&s, b)
	if len(s.argBytes) == 0 {
		t.Fatal("failed stateBeginStringOrArray argBytes length")
	}
	if s.argBytes[0] != 'i' {
		t.Fatal("failed stateBeginStringOrArray argBytes value index 0")
	}

}

func TestStateString1(t *testing.T) {
	var s scanner
	s.initScanner()

	key := "sample"
	for i := 0; i < len(key); i++ {
		s.argBytes = append(s.argBytes, key[i])
	}

	var b byte = '.'

	stateString(&s, b)

	if len(s.result) == 0 {
		t.Fatal("failed TestStateString1 len(s.result)")
	}
	//fmt.Println(s.result)
	if s.result[0].arg != "sample" {
		t.Fatal("failed TestStateString1 s.result[0].arg")
	}
	if s.result[0].isArray == true {
		t.Fatal("failed TestStateString1 s.result[0].isArray")
	}
	if s.result[0].isIndex == true {
		t.Fatal("failed TestStateString1 s.result[0].isIndex")
	}
}

func TestStateString2(t *testing.T) {
	var s scanner
	s.initScanner()

	key := "sample"
	for i := 0; i < len(key); i++ {
		s.argBytes = append(s.argBytes, key[i])
	}

	var b byte = '['

	stateString(&s, b)

	if len(s.result) != 0 {
		t.Fatal("failed TestStateString2 len(s.result)")
	}
}

func TestStateArray(t *testing.T) {
	var s scanner
	s.initScanner()

	key := "sample"
	for i := 0; i < len(key); i++ {
		s.argBytes = append(s.argBytes, key[i])
	}

	var b byte = ']'

	stateArray(&s, b)

	if len(s.result) == 0 {
		t.Fatal("failed TestStateArray len(s.result)")
	}
	if s.result[0].arg != "sample" {
		t.Fatal("failed TestStateArray s.result[0].arg")
	}
	if s.result[0].isArray == false {
		t.Fatal("failed TestStateArray r.result[0].isArray")
	}
}

func TestExtractResult(t *testing.T) {
	var s scanner
	s.initScanner()
	var b byte = '.'
	s.argBytes = append(s.argBytes, b)

	if len(s.argBytes) == 0 {
		t.Fatal("failed TestExtractResult len(s.argsBytes)")
	}
	//fmt.Println(s.argBytes[0])

	s.extractResult()
	if len(s.result) == 0 {
		t.Fatal("failed TestExtractResult len(s.result)")
	}
	//fmt.Println(s.result[0])
}

func TestCreateIndex(t *testing.T) {
	i := "115"
	b := []byte{}
	for j := 0; j < len(i); j++ {
		b = append(b, i[j])
	}

	var s scanner
	s.initScanner()
	s.indexBytes = b

	s.createIndex()
	//fmt.Println(s.index)
	if s.index != 115 {
		t.Fatal("failed TestCreateIndex")
	}
}

func TestStateInt1(t *testing.T) {
	var s scanner
	s.initScanner()
	//s.step = stateInt

	i := "115"
	for j := 0; j < len(i); j++ {
		stateInt(&s, i[j])
	}

	s.createIndex()
	if s.index != 115 {
		t.Fatal("failed TestStateInt1")
	}
}

func TestStateInt2(t *testing.T) {
	var s scanner
	s.initScanner()

	i := "115]"
	for j := 0; j < len(i); j++ {
		stateInt(&s, i[j])
	}

	if len(s.result) == 0 {
		t.Fatal("failed TestStateInt2 len(s.result)")
	}
	if s.result[0].index != 115 {
		t.Fatal("failed TestStateInt2 s.result[0].index")
	}
}
