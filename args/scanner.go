package args

import (
	"strconv"
	"unsafe"
)

//scanResult
//scanResult is object model that each setction of query
//scanResult　クエリの.(ピリオド)ごとにkey名とそれに紐づくデータの構造体を定義したオブジェクト
type scanResult struct {
	//json key name
	//取得するjsonのkey名
	arg string

	//key-value is whether array object or not
	//keyに対するデータが配列かどうか
	isArray bool

	//key-value is whether index has been designated or not
	//keyに対するデータの配列でindexが指定されているかどうか
	isIndex bool

	//array index
	//配列のindexが指定されている場合のindexの数字
	index int
}

//scanner object model for scanning query string
//scanner クエリを読み取るためのオブジェクトモデル
type scanner struct {
	//scan function
	//scan用の関数
	//クエリを1バイトずつ読み取っていく
	step func(*scanner, byte)

	//flag is that confirm whether array object or not
	//セクション（.）ごとのクエリが配列かどうかを示すフラグ
	isArray bool

	//flag is that has index of slice
	//配列だった場合、indexが指定されているかどうかを示すフラグ
	isIndex bool

	//index of slice
	//配列でindexが指定されている場合のindexの数
	index int

	//byte slice before convert to index
	//indexに変換する前のbyteスライス
	indexBytes []byte

	//each section (.) key name
	//セクション(.)ごとのkey名
	middleArg string

	//byte slice before convert to key name
	//middleArgに変換する前のbyteスライス
	argBytes []byte

	//scanResult slice object
	//appent each section (.)
	//セクション（.）ごとのkey-valueのkey名、valueのデータ構造を表したオブジェクトのスライス
	result []scanResult
}

func (s *scanner) initScanner() {
	s.step = stateBeginPeriod
	s.isArray = false
	s.isIndex = false
	s.indexBytes = []byte{}
	s.argBytes = []byte{}
	s.result = []scanResult{}
}

func (s *scanner) resetScanner() {
	s.isArray = false
	s.isIndex = false
	s.argBytes = []byte{}
	s.indexBytes = []byte{}
}

func stateBeginPeriod(s *scanner, b byte) {
	if b == '.' {
		s.step = stateBeginStringOrArray
		return
	}
}

func stateBeginStringOrArray(s *scanner, b byte) {
	if b == '[' {
		s.step = stateArray
		return
	}
	if b != '.' {
		s.step = stateString
		s.argBytes = append(s.argBytes, b)
		return
	}
}

func stateString(s *scanner, b byte) {
	if b == '.' {
		//create scanResult
		s.extractResult()

		//注意: すでに'.'が出てきているので次は [ or string
		s.step = stateBeginStringOrArray
		return
	}
	if b == '[' {
		s.step = stateArray
		return
	}

	s.step = stateString
	s.argBytes = append(s.argBytes, b)
}

func stateArray(s *scanner, b byte) {
	s.isArray = true

	if b == ']' {
		//create scanResult
		s.extractResult()

		s.step = stateBeginPeriod
		return
	}

	s.step = stateInt
	s.indexBytes = append(s.indexBytes, b)
}

func stateInt(s *scanner, b byte) {
	s.isIndex = true

	if b == ']' {
		//create scanResult
		s.extractResult()

		s.step = stateBeginPeriod
		return
	}
	if '0' <= b && b <= '9' {
		s.step = stateInt
		s.indexBytes = append(s.indexBytes, b)
	}
}

//convert to middleArg from byte array(argBytes)
func (s *scanner) createMiddleArg() {
	s.middleArg = *(*string)(unsafe.Pointer(&s.argBytes))
}

//convert to index from byte array(indexBytes)
func (s *scanner) createIndex() {
	value := *(*string)(unsafe.Pointer(&s.indexBytes))
	i, _ := strconv.Atoi(value)
	s.index = i
}

//create scanResult and put in result
func (s *scanner) extractResult() {
	s.createMiddleArg()

	sr := &scanResult{
		arg:     s.middleArg,
		isArray: false,
		isIndex: false,
		index:   0,
	}

	//配列の場合
	if s.isArray == true {
		sr.isArray = true
	}
	//indexありの場合
	if len(s.indexBytes) != 0 {
		sr.isIndex = true
		//indexを入れる
		s.createIndex()
		sr.index = s.index
	}

	s.result = append(s.result, *sr)
	s.resetScanner()
}
