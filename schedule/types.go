package schedule

import (
	"encoding/xml"

	"egbitbucket.dtvops.net/deadline/common"
)

type Schedule struct {
	XMLName  xml.Name       `xml:"schedule"`
	Handler  common.Handler `xml:"handler,omitempty"`
	Timing   string         `xml:"timing,attr,omitempty"`
	Name     string         `xml:"name,attr,omitempty"`
	Schedule []common.Event `xml:"event,omitempty"`
}

type scheduleManager struct {
	subscriptionTable map[string][]Schedule
}

type Error struct {
	To string `xml:"to,attr"`
}

type ScheduleDAO interface {
	GetByName(string) ([]byte, error)
	Save(s Schedule) error
}

type fileDAO struct{}