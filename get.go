package main

// get value

//convert to map[string]interface{} from unmarshal json data (interface{})
func convertMap(i interface{}) (map[string]interface{}, error) {
	m, f := i.(map[string]interface{})
	if f == false {
		err := &Error{
			Message: "cannot convert map[string]interface{}",
		}
		return nil, err
	}
	return m, nil
}

//get value from json key
func getValue(i map[string]interface{}, key string) (interface{}, error) {
	value, f := i[key]
	if f == false {
		err := &Error{
			Message: "no such json key",
		}
		return nil, err
	}
	return value, nil
}

func getValueOfKey(i interface{}, k string) (interface{}, error) {
	m, err := convertMap(i)
	if err != nil {
		return nil, err
	}
	v, err := getValue(m, k)
	if err != nil {
		return nil, err
	}
	return v, nil
}
