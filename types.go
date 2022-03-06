package main

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
)

var ErrNoRecordsFound = errors.New("no records found for that actor")

type Deps struct {
	DB *gocql.Session
}

type Event struct {
	Actor     string     `json:"actor,omitempty"`
	ID        gocql.UUID `json:"id"`
	Type      string     `json:"type"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
}

type EventTimestream struct {
	Actor  string  `json:"actor"`
	Events []Event `json:"events"`
}
