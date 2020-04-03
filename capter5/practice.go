package capter5

import (
	"strconv"
	"strings"
)

func (t *Task) MarkDone() {
	t.Status = DONE
	if t.SubTasks != nil {
		for i := range t.SubTasks {
			(t.SubTasks[i]).MarkDone()
		}
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (si *StructID) UnmarshalJSON(data []byte) error {
	siInt, err := strconv.Atoi(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	*si = StructID(siInt)
	return nil
}

type SortByDeadLineTasks []Task

func (t SortByDeadLineTasks) Len() int { return len(t) }
func (t SortByDeadLineTasks) Less(i, j int) bool {
	if t[i].Deadline == nil || t[j].Deadline == nil {
		return false
	}
	return t[i].Deadline.Before(t[j].Deadline.Time) ||
		t[i].Deadline.Equal(t[j].Deadline.Time)
}
func (t SortByDeadLineTasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type SortByPriorityTasks []Task

func (t SortByPriorityTasks) Len() int { return len(t) }
func (t SortByPriorityTasks) Less(i, j int) bool {
	return t[i].Priority <= t[j].Priority
}
func (t SortByPriorityTasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
