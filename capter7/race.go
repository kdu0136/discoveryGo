package capter7

import (
	"fmt"
	"sync"
)

func RaceMain() {
	SyncOnce()
	//req, resp := make(chan struct{}), make(chan int64)
	//cnt := int64(10)
	//go func(cnt int64) {
	//	defer close(resp)
	//	for range req {
	//		cnt--
	//		resp <- cnt
	//	}
	//}(cnt)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		req <- struct{}{}
	//	}()
	//}
	//for cnt = <-resp; cnt > 0; cnt = <-resp{
	//
	//}
	//close(req)
	//fmt.Println(cnt)
}

func SyncOnceChan() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("initialized")
	}()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-done
			fmt.Println("goroutine:", i)
		}(i)
	}
	wg.Wait()
}

func SyncOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("initialized")
			})
			fmt.Println("goroutine:", i)
		}(i)
	}
	wg.Wait()
}
