package schedule

import (
	"encoding/xml"
	"time"
)

type Schedule struct {
	XMLName  xml.Name       `xml:"schedule"`
	Handler  Handler 		`xml:"handler,omitempty" db:"handler"`
	Timing   string         `xml:"timing,attr,omitempty" db:"timing"`	
	Name     string         `xml:"name,attr,omitempty" db:"name"`
	Schedule []byte         `xml:",innerxml"`
	LastRun	 time.Time
	Start    Node           `xml:"-"`
	End      Node           `xml:"-"`
	Error    Node           `xml:"-"`
}

type Event struct {
	XMLName   xml.Name          `xml:"event"`
	Name      string            `json:"name" xml:"name,attr" db:"name"`
	Success   bool              `json:"success" xml:"success" db:"success"`
	Details   map[string]string `json:"details,omitempty" xml:"details,omitempty" db:"details"`
	ReceiveBy string            `xml:"receive-by,attr" db:"receiveby"`
	ReceiveAt string            `xml:"receive-at,attr" db:"receiveat"`
	IsLive bool `xml:"islive"`
}

type ScheduledEvent struct {
	ScheduleName 	string	`db:"schedulename"`
	EName			string  `db:"ename"` 
	EReceiveBy		string  `db:"ereceiveby"` 

}

type Handler struct {
	XMLName xml.Name `xml:"handler"`
	Name    string   `xml:"name,attr" db:"name"`
	Address string   `xml:"address"   db:"address"`
}

type ScheduledHandler struct {
	ScheduleName 	string	`db:"schedulename"`
	Name			string  `db:"name"` 
	Address			string  `db:"address"` 

}

type Node struct {
	Event   *Event `xml:"event"`
	Nodes   []Node        `xml:"-"`
	ErrorTo *Node         `xml:"-"`
	OkTo    *Node         `xml:"-"`
}

type Start struct {
	Node
}

type End struct {
	Node
}

type ScheduleManager struct {
	subscriptionTable map[string][]*Schedule
	EvaluationTime	time.Time
}

type Error struct {
	To string `xml:"to,attr"`
}

type ScheduleDAO interface {
	GetByName(string) ([]byte, error)
	Save(s *Schedule) error
	LoadStatelessSchedules() ([]Schedule,error)
	LoadEvents() ([]Event,error)
	SaveEvent( e *Event) error
}

type fileDAO struct{
	Path string

}

type dbDAO struct {
	ConnectionString string
}

type innerbytes struct {
	XMLName xml.Name `xml:"innerbytes"`
	Hander Handler 	`xml:"handler"`
	Events []Event	`xml:"event"`
}