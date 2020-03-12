package capter3

import "fmt"

func ExampleHasConsonantSuffixHangul() {
	fmt.Println(HasConsonantSuffixHangul("위"))
	fmt.Println(HasConsonantSuffixHangul("컴"))
	fmt.Println(HasConsonantSuffixHangul("안"))
	// Output:
	// false
	// true
	// true
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
}