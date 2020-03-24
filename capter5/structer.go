package capter5

import (
	"errors"
	"strconv"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return "unknown value"
	}
}

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

// OverDue returns true if the deadline is before the current time.
func (d *Deadline) OverDue() bool {
	return d != nil && d.Time.Before(time.Now())
}

// OverDue returns true if the deadline is before the current time.
func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}

type MyStruct struct {
	Title    string `json:"title"`
	Internal string `json:"-"`          // json 무시
	Value    int64  `json:",omitempty"` // 0일경우 무시
	ID       int64  `json:",string"`    // json 에서 문자
}

// MarshalJON implements the json.Marshaler interface
func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}
