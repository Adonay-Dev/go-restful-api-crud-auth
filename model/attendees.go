package model

import (
	"github.com/google/uuid"
)

type attendees struct {
	Id      uuid.UUID `json:"id"`
	UserId  User      `json:"userid"`
	EventId Event     `json:"eventid"`
}
