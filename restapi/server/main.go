package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/puslip41/GoStudy/restapi"
)

var storage restapi.MemoryStorage

const uriHeader = "/users/"

func main() {
	http.HandleFunc(uriHeader, UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "GET":
		id := getMemberID(r.URL.Path)
		if name, email, err := storage.Query(id); err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusFound)
			fmt.Fprintf(w, `{"id":"%s","name":"%s","email":"%s"}`, id, name, email)
		}
		break

	case "POST":
		id := getMemberID(r.URL.Path)

		if password := r.FormValue("password"); len(password) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else if name:= r.FormValue("name"); len(name) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else if email:= r.FormValue("email"); len(email) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else {
			if err := storage.Register(id, password, name, email); err != nil {
				w.WriteHeader(http.StatusNotAcceptable)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		}
		break

	case "PUT":
		id := getMemberID(r.URL.Path)

		if password := r.FormValue("password"); len(password) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else if name:= r.FormValue("name"); len(name) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else if email:= r.FormValue("email"); len(email) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else {
			if err := storage.Modify(id, password, name, email); err != nil {
				w.WriteHeader(http.StatusNotModified)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusAccepted)
			}
		}
		break

	case "DELETE":
		id := getMemberID(r.URL.Path)

		if password := r.FormValue("password"); len(password) == 0 {
			w.WriteHeader(http.StatusPreconditionRequired)
		} else {
			if err := storage.Delete(id, password); err != nil {
				w.WriteHeader(http.StatusNotAcceptable)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusAccepted)
			}
		}
		break

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getMemberID(path string) string {
	return path[len(uriHeader):]
}
