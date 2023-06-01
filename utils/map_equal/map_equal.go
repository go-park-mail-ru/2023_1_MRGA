package map_equal

import "fmt"

func MapEqual(got, expected map[string]interface{}) bool {
	for keyGot, valueGot := range got {
		valueExpected := expected[keyGot]

		var (
			strValueExpected string = convertToString(valueExpected)
			strValueGot      string = convertToString(valueGot)
		)
		if strValueExpected != strValueGot {
			return false
		}
	}
	return true
}

func convertToString(val interface{}) (strVal string) {
	switch v := val.(type) {
	case string:
		strVal = v
	case float64:
		strVal = fmt.Sprint(v)
	case int:
		strVal = fmt.Sprint(v)
	}
	return strVal
}
