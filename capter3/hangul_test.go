package capter3

import "fmt"

func ExampleHasConsonantSuffixHangul() {
	fmt.Println(HasConsonantSuffixHangul("사과"))
	fmt.Println(HasConsonantSuffixHangul("수박"))
	fmt.Println(HasConsonantSuffixHangul("바나나"))
	// Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"
	for i:=0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	// Output:
	// 234
	// 176
	// 128
	// 235
	// 130
	// 152
	// 235
	// 139
	// 164
}