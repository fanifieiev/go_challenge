package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-challenge/http/server/model"
	"github.com/gorilla/mux"
)

// HTTPGetNote Get handler
func HTTPGetNote(rw http.ResponseWriter, rq *http.Request) {
	noteID := mux.Vars(rq)["noteId"]
	i, err := strconv.Atoi(noteID)
	if err != nil {
		panic(err)
	}

	var note *model.Note = &model.Note{
		ID:       int64(i),
		Title:    "Fevzi",
		Context:  "Context",
		Author:   "Nobody",
		IsPublic: true,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(note); err != nil {
		http.Error(rw, "Failed to encode note", http.StatusInternalServerError)
		return
	}

}
