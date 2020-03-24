package capter4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rx = regexp.MustCompile(`{[^}]+}`)

func EvalReplaceAll(in string) {
	eval := NewEvaluator()
	replacer := strings.NewReplacer("{", "", "}", "")

	out := rx.ReplaceAllStringFunc(in, func(expr string) string {
		expr = replacer.Replace(expr)
		num, _ := eval(expr)
		return strconv.Itoa(num)
	})
	fmt.Println(out)
}
