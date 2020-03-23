package capter5

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func ExampleDeadline_OverDue() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1}
	t2 := Task{"4h later", TODO, d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

func Example_marshalJSON() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	t2 := MyStruct{
		Title:    "my title",
		Internal: "ignore",
		Value:    123,
		ID:       1234,
	}
	b2, err := json.Marshal(t2)
	fmt.Println(string(b2))
	// Output:
	// {"Title":"Laundry","Status":"DONE","Deadline":1439739780}
	// {"title":"my title","Value":123,"ID":"1234"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Laundry","Status":"DONE", "Deadline":1439739780}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t)
	b2 := []byte(`{"title":"my title", "Internal":"ignore", "value":2, "ID":"1234"}`)
	t2 := MyStruct{}
	err = json.Unmarshal(b2, &t2)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t2)
	// Output:
	// {Laundry DONE 2015-08-17 00:43:00 +0900 KST}
	// {my title  2 1234}
}
