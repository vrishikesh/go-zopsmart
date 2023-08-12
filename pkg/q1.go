package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// create a server with GET /users and return response {data:{users:[{id:1,name:rishi}]}}
// if /users receives any other http method then return status 404 and text method not allowed
// and for any other path just return status 404
type HTTPResponse struct {
	Data HTTPResponseData `json:"data"`
}

type HTTPResponseData struct {
	Users []User `json:"users"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func Q1() {
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: http.HandlerFunc(HandleRequest),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: method [%s], path [%s]\n", r.Method, r.URL.Path)

	if r.URL.Path != "/users" {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("response: status [%d]\n", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		content := "method not allowed"
		log.Printf("response: status [%d], content: [%s]\n", http.StatusMethodNotAllowed, content)
		fmt.Fprint(w, content)
		return
	}

	u := User{ID: 1, Name: "rishi"}
	d := HTTPResponseData{Users: []User{u}}
	res := HTTPResponse{Data: d}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	content := string(b)
	log.Printf("response: status [%d], content: [%s]\n", http.StatusOK, content)
	fmt.Fprint(w, content)
}
