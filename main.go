package main

import (
	"encoding/json"
	"errors"
	"github.com/discoveryGo/capter5"
	"github.com/discoveryGo/capter6"
	"log"
	"net/http"
)

// FIXME: m is NOT thread-safe
var m = capter6.NewMemoryDataAccess()

const pathPrefix = "/api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (capter6.ID, error) {
		id := capter6.ID(r.URL.Path[len(pathPrefix):])
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
		err = json.NewEncoder(w).Encode(capter6.Response{
			ID:    id,
			Task:  t,
			Error: capter6.ResponseError{Err: err},
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
			err = json.NewEncoder(w).Encode(capter6.Response{
				ID:    id,
				Task:  t,
				Error: capter6.ResponseError{Err: err},
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
			err = json.NewEncoder(w).Encode(capter6.Response{
				ID:    id,
				Task:  t,
				Error: capter6.ResponseError{Err: err},
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
		err = json.NewEncoder(w).Encode(capter6.Response{
			ID:    id,
			Error: capter6.ResponseError{Err: err},
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
