package main

import (
	"github.com/discoveryGo/capter4"
	"strings"
)

func main() {
	in := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 { 10 ** 5 mod 7777 }평을 가져라.",
		"막내는 { 10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고",
		"배기랑 { 711 * 8 / 9 }cc의 경운기를 갖거라",
	}, "\n")

	capter4.EvalReplaceAll(in)
}
