package capter6

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/discoveryGo/capter5"
)

// ID is data type to identify a task
type ID string

// DataAccess is an interface to access task
type DataAccess interface {
	Get(id ID) (capter5.Task, error)
	Put(id ID, t capter5.Task) error
	Post(t capter5.Task) (ID, error)
	Delete(id ID) error
}

// MemoryDataAccess is a simple in-memory database.
type MemoryDataAccess struct {
	tasks  map[ID]capter5.Task
	nextID int64
}

// ErrTaskNotExist occurs when the task with the ID was not found.
var ErrTaskNotExist = errors.New("task does not exist")

// Get returns a task with as given ID.
func (m *MemoryDataAccess) Get(id ID) (capter5.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return capter5.Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put updates a task with a given ID with t.
func (m *MemoryDataAccess) Put(id ID, t capter5.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new task.
func (m *MemoryDataAccess) Post(t capter5.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID.
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

// NewMemoryDataAccess returns a new MemoryDataAccess.
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]capter5.Task{},
		nextID: int64(1),
	}
}

// ResponseError is the error for the JSON Response.
type ResponseError struct {
	Err error
}

// MarshalJSON returns the JSON representation of the error.
func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

// UnmarshalJSON parses the JSON representation of the error
func (err *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNotExist.Error() {
			err.Err = ErrTaskNotExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

// Response is a struct for the JSON response.
type Response struct {
	ID    ID            `json:"id,omitempty"`
	Task  capter5.Task  `json:"task"`
	Error ResponseError `json:"error"`
}
