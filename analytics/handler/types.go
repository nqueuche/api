package handler

import "time"
import "github.com/segmentio/ksuid"

type User struct {
	ID        ksuid.KSUID
	Username  string
	Email     string
	CreatedAt time.Time
}

type Events struct {
	ID         ksuid.KSUID
	Name       string
	User_id    string
	Library_id string
}

type Library struct {
	ID      ksuid.KSUID
	Name    string
	Version string
}
