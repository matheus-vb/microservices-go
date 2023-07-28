package main

import (
	"net/http"

	"github.com/matheus-vb/microservices-go/logger-service/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	event := data.LogEntry{
		Data: requestPayload.Data,
		Name: requestPayload.Name,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Logged.",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
