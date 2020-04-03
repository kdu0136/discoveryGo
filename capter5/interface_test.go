package capter5

import (
	"container/heap"
	"fmt"
	"sort"
)

func ExampleTask_String() {
	fmt.Println(Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
	})
	// Output:
	// [v] Laundry <nil>
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
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
	}))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Put <nil>
	//     [ ] Detergent <nil>
	//   [ ] Dry <nil>
	//   [ ] Fold <nil>
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	fmt.Println(apple)
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [iPhone iPad MacBook AppStore]
	// [AppStore iPad iPhone MacBook]
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

func ExampleCaseInsensitive_heapString() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}
