package args

/*

.
.[]
.item[]
.item[3]

*/

type jsonState struct {
	//interface object after unmarshal json (input)
	//data interface{}
	//slice of scanResult
	queryResult []scanResult
	//Data of output read from queryResult
	statement interface{}
	//
	isArray bool
}

func (js *jsonState) jsonread() {
	for _, sr := range js.queryResult {
		js.argCheck(sr)
	}
}

func (js *jsonState) argCheck(sr scanResult) {
	if sr.arg != "" {
		if js.isArray == false {
			//map[string]interfaceに変換してkeyで取る
			js.extractObjectFromKey(sr.arg)
		} else {
			js.scanStatement(sr.arg)
		}
	}
	//配列かどうか調べる
	js.arrayCheck(sr)
}

func (js *jsonState) arrayCheck(sr scanResult) {
	if sr.isArray == false {
		return
	}
	//
	r := js.extractArray()

	if sr.isIndex == false {
		js.statement = interface{}(r)
		js.isArray = true
		return
	}

	if sr.index > len(r) {
		return
	}
	js.statement = r[sr.index]
}

func (js *jsonState) extractObjectFromKey(arg string) {
	r, _ := js.statement.(map[string]interface{})[arg]
	js.statement = r
}

func (js *jsonState) extractArray() []interface{} {
	r, _ := js.statement.([]interface{})
	return r
}

func (js *jsonState) scanStatement(arg string) {
	r, _ := js.statement.([]interface{})
	n := []interface{}{}
	for _, v := range r {
		val, b := v.(map[string]interface{})
		if b == true {
			n = append(n, val[arg])
		} else {
			scanArrayStatement(&n, v, arg)
		}
	}
	js.statement = interface{}(n)
}

func scanArrayStatement(n *[]interface{}, v interface{}, arg string) {
	r, _ := v.([]interface{})
	for _, cr := range r {
		gcr, _ := cr.(map[string]interface{})[arg]
		*n = append(*n, gcr)
	}
}
