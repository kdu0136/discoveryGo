package capter6

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/discoveryGo/capter5"
	"html/template"
	"log"
	"net/http"
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

// FIXME: m is NOT thread-safe
var m = NewMemoryDataAccess()

const PathPrefix = "/api/v1/task/"

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(PathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}

	getTasks := func() ([]capter5.Task, error) {
		var result []capter5.Task
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			result = append(result, capter5.Task{
				Title: encodedTask,
			})
		}
		return result, nil
	}

	switch r.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{Err: err},
		})
		if err != nil {
			log.Println(err)
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{Err: err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "POST":
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{Err: err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Error: ResponseError{Err: err},
		})
		if err != nil {
			log.Println(err)
		}
	}
}

const HtmlPrefix = "/task/"

var tmpl = template.Must(template.ParseGlob("html/*.html"))

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(r.Method, "method is not supported")
		return
	}
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(HtmlPrefix):])
		if id == "" {
			return id, errors.New("htmlHandler: ID is empty")
		}
		return id, nil
	}
	id, err := getID()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := m.Get(id)
	err = tmpl.ExecuteTemplate(w, "task.html", &Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
