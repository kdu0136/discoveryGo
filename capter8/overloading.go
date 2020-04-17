package capter8

import "fmt"

type Option struct {
	Idx int
	Lang string
	ExcludeEmpty bool
}

type Element struct {
	opt Option
}

func GetElement(opt Option) *Element {
	return &Element{opt:opt}
}

type Stringer interface {
	String() string
}

type Int int
type Double float64

func (i Int) String() string { return fmt.Sprintf("%d", i) }
func (d Double) String() string { return fmt.Sprintf("%f", d) }
