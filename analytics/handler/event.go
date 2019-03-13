package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/segmentio/ksuid"
)

type Event struct {
	ID        ksuid.KSUID
	Name      string
	User      *User
	Library   *Library
	CreatedAt time.Time
}

type PayloadEvent struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func postEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var body ResponsePostEvent

	var event PayloadEvent
	err := decoder.Decode(&event)
	if err != nil {
		body = ResponsePostEvent{
			StatusCode: 400,
			Message:    "Bad data",
			Data:       nil,
		}

		bytes, _ := json.Marshal(body)
		w.WriteHeader(400)
		w.Write(bytes)
		return
	}

	db, err := NewDatabase()
	defer db.Close()

	if err != nil {
		body = ResponsePostEvent{
			StatusCode: 500,
			Message:    "Internal server error",
			Data:       nil,
		}

		bytes, _ := json.Marshal(body)

		w.WriteHeader(500)
		w.Write(bytes)
		return
	}

	id := ksuid.New()

	_, err = db.Query("INSERT INTO events (id, name, user_id) VALUES ($1, $2, $3)", id, event.Name, event.UserID)
	if err != nil {
		body = ResponsePostEvent{
			StatusCode: 400,
			Message:    "Bad data",
			Data:       nil,
		}

		bytes, _ := json.Marshal(body)

		w.WriteHeader(400)
		w.Write(bytes)
		return
	}

	body = ResponsePostEvent{
		StatusCode: 201,
		Message:    "Created",
		Data: &ResponsePostEventData{
			Name: "arthur",
		},
	}
	bytes, _ := json.Marshal(body)

	w.WriteHeader(201)
	w.Write(bytes)
}
