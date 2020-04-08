package capter7

import "fmt"

func Example_simpleChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}
	for num := range c() {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, " ")
	}
	// Output:
	// 0 1 1 2 3 5 8 13
}

func ExampleCreateName() {
	for name := range CreateName("성정", "준호") {
		fmt.Println(name)
	}
	// Output:
	// 성준
	// 성호
	// 정준
	// 정호
}

func ExampleClosedChannel() {
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 0
	// 0
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(PlusOne(c))) {
		fmt.Println(num)
	}
	// Output:
}

func ExampleChain() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range Chain(PlusOne, PlusOne, PlusOne)(c) {
		fmt.Println(num)
	}
	// Output:
}
