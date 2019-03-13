package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/segmentio/ksuid"
)

type User struct {
	ID        ksuid.KSUID
	Username  string
	Email     string
	CreatedAt time.Time
}

type PayloadUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func postUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var user PayloadUser
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	db, err := NewDatabase()
	defer db.Close()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	id := ksuid.New()

	_, err = db.Query("INSERT INTO users (id, username, email) VALUES ($1, $2, $3)", id.String(), user.Username, user.Email)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := []byte(`{ "hello": "world" }`)
	w.WriteHeader(201)
	w.Write(body)
}
