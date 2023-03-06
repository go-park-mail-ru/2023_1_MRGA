package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

func Checkout(ts *httptest.Server, inputJson string) (result map[string]interface{}, err error) {
	var resp *http.Response
	resp, err = http.Post(ts.URL, "application/json", bytes.NewBuffer([]byte(inputJson)))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var jsonStr []byte
	jsonStr, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func convertToString(val interface{}) (strVal string) {
	switch val.(type) {
	case string:
		strVal = val.(string)
	case float64:
		strVal = fmt.Sprint(val.(float64))
	case int:
		strVal = fmt.Sprint(val.(int))
	}
	return strVal
}

func mapEqual(got, expected map[string]interface{}) bool {
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

func TestRegister(t *testing.T) {
	// tests := map[string]struct {
	// 	inputJson  string
	// 	outputJson map[string]interface{}
	// }{
	// 	"Обычный тест на создание пользователя": {
	// 		inputJson:  `{"username": "masharpik", "email": "masharpik@gmail.com", "password": "masharpik2004", "age": 19}`,
	// 		outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
	// 	},
	// 	"Создание пользователя с тем же ником": {
	// 		inputJson:  `{"username": "masharpik", "email": "masharpik2@gmail.com", "password": "masharpik2004", "age": 19}`,
	// 		outputJson: map[string]interface{}{"err": "username is not unique", "status": http.StatusBadRequest},
	// 	},
	// }

	// r := repository.NewRepo()
	// a := New(r)
	// ts := httptest.NewServer(http.HandlerFunc(a.Register))

	// for testName, test := range tests {
	// 	testName := testName
	// 	test := test
	// 	t.Run(testName, func(t *testing.T) {
	// 		result, err := Checkout(ts, test.inputJson)
	// 		if err != nil {
	// 			t.Errorf("[%s] unexpected error: %#v", testName, err)
	// 		}
	// 		if !mapEqual(result, test.outputJson) {
	// 			t.Errorf("[%s] wrong result, expected %#v, got %#v", testName, test.outputJson, result)
	// 		}
	// 	})
	// }
}
