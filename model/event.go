package model

import (
	"github.com/google/uuid"
)

type Event struct {
	Id               uuid.UUID `json:"id"`
	OwnerId          string    `json:"ownerid"`
	EventName        string    `json:"eventname"`
	EventDeScription string    `json:"eventdescription"`
	EventDate        string    `json:"eventdate"`
}
