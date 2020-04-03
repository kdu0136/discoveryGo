package capter5

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

func ExampleTask_MarkDone() {
	t := Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{{
				Title:    "Put",
				Status:   DONE,
				Deadline: nil,
				Priority: 2,
			}, {
				Title:    "Detergent",
				Status:   TODO,
				Deadline: nil,
				Priority: 2,
			}},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
		}},
	}
	(&t).MarkDone()
	fmt.Println(IncludeSubTasks(t))
	// Output:
	// [v] Laundry <nil>
	//   [v] Wash <nil>
	//     [v] Put <nil>
	//     [v] Detergent <nil>
	//   [v] Dry <nil>
	//   [v] Fold <nil>
}

func ExampleStructID_UnmarshalJSON() {
	b := []byte(`{"title":"my title", "Internal":"ignore", "value":2, "id":"1234"}`)
	t := MyStruct{}
	if err := json.Unmarshal(b, &t); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	b2 := []byte(`{"title":"my title", "Internal":"ignore", "value":2, "id":1234}`)
	t2 := MyStruct{}
	if err := json.Unmarshal(b2, &t2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t2)
	// Output:
	// {my title  2 1234}
	// {my title  2 1234}
}

func ExampleTasks_sort() {
	t := []Task{
		{
			Title:    "Priority4",
			Status:   TODO,
			Deadline: NewDeadline(time.Date(2019, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 4,
		},
		{
			Title:    "Priority2",
			Status:   TODO,
			Deadline: NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 2,
		},
		{
			Title:    "Priority3",
			Status:   TODO,
			Deadline: NewDeadline(time.Date(2017, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 3,
		},
		{
			Title:    "Priority1",
			Status:   TODO,
			Deadline: NewDeadline(time.Date(2020, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 1,
		},
		{
			Title:    "Priority5",
			Status:   TODO,
			Deadline: NewDeadline(time.Date(2018, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 5,
		},
	}
	st := SortByDeadLineTasks(t)
	fmt.Println(st)
	sort.Sort(st)
	fmt.Println(st)
	st2 := SortByPriorityTasks(t)
	sort.Sort(st2)
	fmt.Println(st2)
	// Output:
	// [[ ] Priority4 2019-08-16 15:43:00 +0000 UTC [ ] Priority2 2015-08-16 15:43:00 +0000 UTC [ ] Priority3 2017-08-16 15:43:00 +0000 UTC [ ] Priority1 2020-08-16 15:43:00 +0000 UTC [ ] Priority5 2018-08-16 15:43:00 +0000 UTC]
	// [[ ] Priority2 2015-08-16 15:43:00 +0000 UTC [ ] Priority3 2017-08-16 15:43:00 +0000 UTC [ ] Priority5 2018-08-16 15:43:00 +0000 UTC [ ] Priority4 2019-08-16 15:43:00 +0000 UTC [ ] Priority1 2020-08-16 15:43:00 +0000 UTC]
	// [[ ] Priority1 2020-08-16 15:43:00 +0000 UTC [ ] Priority2 2015-08-16 15:43:00 +0000 UTC [ ] Priority3 2017-08-16 15:43:00 +0000 UTC [ ] Priority4 2019-08-16 15:43:00 +0000 UTC [ ] Priority5 2018-08-16 15:43:00 +0000 UTC]
}
