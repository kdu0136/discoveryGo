package capter5

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Loader interface {
	Load(filename string) error
}

type ReadWriter interface {
	io.Reader
	io.Writer
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func PrintStringer(data fmt.Stringer) {
	fmt.Print(data.String())
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

type CaseInsensitive []string

func (c CaseInsensitive) Len() int { return len(c) }
func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}
func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}
func (c *CaseInsensitive) Pop() interface{} {
	cLen := c.Len()
	last := (*c)[cLen-1]
	*c = (*c)[:cLen-1]
	return last
}

type FileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}

func ManageFiles(fs FileSystem) {
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range t {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}
