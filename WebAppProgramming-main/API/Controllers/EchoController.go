package Controllers

import (
	"fmt"
	ioutil "io/ioutil"
	"net/http"
)

func GetEchoController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World Controller")
	w.Write([]byte("hello world!!!"))
}

func PostEchoController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	w.Write(body)
}
