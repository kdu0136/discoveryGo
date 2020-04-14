package capter7

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ChannelMain() {
	max := 100
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Print(prime, " ")
	}
	fmt.Println()
	//reqs := make(chan Request)
	//defer close(reqs)
	//for i := 0; i < 3; i++ {
	//	go PlusOneService(reqs, i)
	//}
	//var wg sync.WaitGroup
	//for i := 3; i < 53; i += 10 {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		resps := make(chan Response)
	//		reqs <- Request{i, resps}
	//		//fmt.Println(i, "=>", <-resps)
	//		for resp := range resps {
	//			fmt.Println(i, "=>", resp)
	//		}
	//	}(i)
	//}
	//wg.Wait()

	//c := make(chan int)
	//go func() {
	//	defer close(c)
	//	for i := 3; i < 103; i += 10 {
	//		c <- i
	//	}
	//}()
	//ctx, cancel := context.WithCancel(context.Background())
	//nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	//for num := range nums {
	//	fmt.Println(num)
	//	if num >= 18 {
	//		cancel()
	//		break
	//	}
	//}
	//time.Sleep(100 * time.Millisecond)
	//fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	//for _ = range nums {
	//	// Consume all nums
	//}
	//time.Sleep(100 * time.Millisecond)
	//fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	//c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	//sendInts := func(c chan<-int, begin, end int) {
	//	defer close(c)
	//	for i := begin; i < end; i++ {
	//		time.Sleep(500 * time.Millisecond)
	//		c <- i
	//	}
	//}
	//go sendInts(c1, 11, 14)
	//go sendInts(c2, 21, 23)
	//go sendInts(c3, 31, 35)
	//for n := range FanIn3(c1, c2, c3) {
	//	fmt.Println(n)
	//}
}

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func CreateName(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

// PlusOne returns a channel of num+1 for nums received from in.
// When done channel is closed, the output channel is closed as well.
func PlusOne(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func FanOut() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		timeout := time.After(1500 * time.Millisecond)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			case <-timeout:
				fmt.Println("time out")
				return
			}
		}
	}()
	return out
}

type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

// Range returns a channel and sends ints
// (start, start+step, start+2*step, ...).
func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

type IntPipe2 func(context.Context, <-chan int) <-chan int

// FilterMultiple returns a IntPipe the filters multiple of n.
func FilterMultiple(n int) IntPipe2 {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

// Primes returns prime number until context done.
func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
