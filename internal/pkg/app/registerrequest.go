package app

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/utils"
)

func (a *Application) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error method", http.StatusBadRequest)

		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error cant read json", http.StatusBadRequest)

		return

	}

	var userJson ds.User
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {

		http.Error(w, "error cant parse json", http.StatusBadRequest)

		return
	}

	hashedPass := CreatePass(userJson.Password)
	userJson.Password = hashedPass

	err = a.repo.AddUser(&userJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	m := utils.Message(true, "success")
	utils.Respond(w, m)
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
	return string(bs)
}
