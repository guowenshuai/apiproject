package models

import (
	"container/list"
)

type EventType int

const (
	EVENT_MESSAGE = iota
	EVENT_JOIN
	EVENT_LEAVE
)

type Event struct {
	Type      EventType
	Uid       string
	Timestamp int
	Content   string
}

const archiveSize = 20

var archive = list.New()

// NewEvent add a New Event to archive list
func NewEvent(e Event) {
	if archive.Len() > archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(e)
}


// GetEvent return all event
func GetEvent() (events []Event) {
	//events := make([]Event, 0, archive.Len())
	for e := archive.Front(); e != nil; e = e.Next() {
		v := e.Value.(Event)
		events = append(events, v)
	}
	return
}
