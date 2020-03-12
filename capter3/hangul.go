package capter3

import "fmt"

func HangulUnicode() {
	// i - 각 유니코드 문자에 대한 바이트 위치(index) int = 3
	// r - 각 문자 rune
	for i, r := range "가갛힣" {
		fmt.Println(i, r)
	}
}

// HasConsonantSuffixHangul returns true f s has hangul consonant at the end
func HasConsonantSuffixHangul(s string) (result bool) {
	start := rune(44032) // "가"의 유니코드 포인트
	end := rune(55204) // "힣"의 유니코드 포인트

	numEnds := 28
	for _, r := range s {
		if start <= r && r < end { // 한글 체크
			index := int(r-start)
			result = index%numEnds != 0 // 0 이면 받침 x
		}
	}
	return
}
