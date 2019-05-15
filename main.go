package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"golang-simple-bank/model"
)

var (
	storage = model.NewMemStorage()
)

func HttpOK(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, msg)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	id := &model.Id{}
	if err := decoder.Decode(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := storage.GetAccount(id.Account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HttpOK(w, acc.String())
}

func AddAccount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	acc := &model.Account{}
	if err := decoder.Decode(acc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := storage.AddAccount(*acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HttpOK(w, acc.String())
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	tr := &model.Transfer{}
	if err := decoder.Decode(tr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := storage.Transfer(tr.Sender, tr.Receiver, tr.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HttpOK(w, tr.String())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get_account", GetAccount)
	r.HandleFunc("/add_account", AddAccount)
	r.HandleFunc("/transfer", Transfer)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
