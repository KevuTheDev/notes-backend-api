package main

import (
	"fmt"
	"net/http"
)

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"ping": "pong",
		"foo":  "bar",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
