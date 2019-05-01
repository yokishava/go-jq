package args

//ArgState have output data and input data, query...
type ArgState struct {
	//json data
	data interface{}
	//query string
	arg string
	//scanner
	scan scanner
	//output data
	result interface{}
}

//ExtractValue extract value from json data
func ExtractValue(d interface{}, a string, r *interface{}) {
	var as ArgState
	as.init(d, a, r)
	as.argWhile()
}

//initialize ArgState
func (as *ArgState) init(d interface{}, a string, r *interface{}) *ArgState {
	as.data = d
	as.arg = a
	as.result = r
	as.scan.initScanner()
	return as
}

//check arg (query), one character at a time
func (as *ArgState) argWhile() {
	str, s := as.arg, &as.scan
	for i := 0; i < len(str); i++ {
		s.step(s, str[i])
	}
	if len(s.argBytes) != 0 {
		s.extractResult()
		return
	}
	if s.isArray == true {
		s.extractResult()
		return
	}
}
