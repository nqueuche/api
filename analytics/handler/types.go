package handler

import (
	"github.com/segmentio/ksuid"
)

type Library struct {
	ID      ksuid.KSUID
	Name    string
	Version string
}
