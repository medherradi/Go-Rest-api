package models

import "time"

type Event struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

var events = []Event{} // slice of event

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
