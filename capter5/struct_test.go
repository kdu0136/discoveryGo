package capter5

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func ExampleDeadline_OverDue() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{
		Title:    "4h ago",
		Status:   TODO,
		Deadline: d1,
	}
	t2 := Task{
		Title:    "4h later",
		Status:   TODO,
		Deadline: d2,
	}
	t3 := Task{
		Title:    "no due",
		Status:   TODO,
		Deadline: nil,
	}
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
	// {"title":"Laundry","status":"DONE","deadline":1439739780}
	// {"title":"my title","Value":123,"id":1234}
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
	// [v] Laundry 2015-08-17 00:43:00 +0900 KST
	// {my title  2 1234}
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"`
		Additional     string `json:"additional,omitempty"`
	}{
		Fields:     f,
		Additional: "c",
	})
	fmt.Println(string(b))
	// Output:
	// {"visibleField":"a","additional":"c"}
}

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width int = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
	// Output:
	// 0e ff 81 04 01 02 ff 82 00 01 0c 01 0c 00 00 08
	// ff 82 00 01 01 4e 01 4a
	// map[N:J]
}
