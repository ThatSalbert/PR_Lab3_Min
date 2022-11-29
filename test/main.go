package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"http_server/items"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()

	go items.GetInstance()

	items.GetInstance().Create(12311, "Test")
	items.GetInstance().Create(99231, "Test2")

	router.HandleFunc("/datastore", ReadAll).Methods("GET")
	router.HandleFunc("/datastore/get", GetByID).Methods("GET")
	router.HandleFunc("/datastore/create", PostByID).Methods("POST")
	router.HandleFunc("/datastore/update", UpdateByID).Methods("PUT")
	router.HandleFunc("/datastore/remove", RemoveByID).Methods("DELETE")

	err := http.ListenAndServe(":8001", router)
	if err != nil {
		return
	}
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ReadAll called.")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items.GetInstance())
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID called.")
	UserId, _ := strconv.Atoi(r.URL.Query().Get("userid"))

	if items.GetInstance().Check(UserId) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(items.GetInstance().Read(UserId))
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error."))
		return
	}
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID called.")
	UserId, _ := strconv.Atoi(r.URL.Query().Get("userid"))
	Message := r.URL.Query().Get("message")

	if items.GetInstance().Check(UserId) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(items.GetInstance().Update(UserId, Message))
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error."))
		return
	}
}

func RemoveByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemoveByID called.")
	UserId, _ := strconv.Atoi(r.URL.Query().Get("userid"))

	if items.GetInstance().Check(UserId) {
		items.GetInstance().Delete(UserId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted user: " + strconv.Itoa(UserId)))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error."))
		return
	}
}

func PostByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostByID called.")
	UserId, _ := strconv.Atoi(r.URL.Query().Get("userid"))
	Message := r.URL.Query().Get("message")

	if items.GetInstance().Check(UserId) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Error."))
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		data := items.GetInstance().Create(UserId, Message)
		res, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
