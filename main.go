package main

import (
	"context"
	"fmt"
	"time"
)

func ex1() {
	time1 := time.Now()

	for i := 0; i < 3; i++ {

		fmt.Println("time now: ", time1)
		if i == 2 {
			fmt.Println("ket thup")
		}
		time.Sleep(3 * time.Second)
	}
}

func ex2() {
	timeq := time.Now().Unix()
	fmt.Println("Timestap hiện tại: ", timeq)
}

func text(ctx context.Context) {
	fmt.Println("text")
	time.Sleep(1 * time.Second)
}

func ex3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 3; i++ {
		fmt.Println("sleep: ", i+1)
		text(ctx)
		time.Sleep(3 * time.Second)
	}
}

func ex4() {
	time1 := time.UnixMilli(1592190294764144364)
	h, _ := time.ParseDuration(time1.String())
	fmt.Println("Chuyển sang thời gian: ", time1)
	fmt.Printf("Số phút: %f\n", h.Minutes())
}

func ex5() {
	a := 1592190385
	now := time.Unix(int64(a), 0)
	fmt.Println(now)
	t := now
	fmt.Println(t.String())
}

func ex6() {
	fmt.Println("Các mốc đơn vị: date, unix")
}

func ex7() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	timenow := time.Now().Unix()
	fmt.Println("Thời gian hiện tại:", timenow)
	go x(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Thời gian sau 3s: ", x(ctx))
	hieu := x(ctx) - timenow
	println("Hiệu thời gian sau 3s:", hieu)
}

func x(ctx context.Context) int64 {
	time1 := time.Now().Unix()
	return int64(time1)
}

func ex8() {
	timenow := "${time.Now().Unix()}"
	for n := 0; n < 5; {
		fmt.Println("Time: ", timenow)
		time.Sleep(100 * time.Millisecond)
		n++

	}
}

func ex9() {
	ch := make(chan int)
	time.AfterFunc(100*time.Millisecond, func() {

		ch <- 10
	})
	for {
		select {
		case i := <-ch:
			fmt.Println("i'm study", i)
			return
		default:
			fmt.Println("wait")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ex1()
	fmt.Println("==========")
	ex2()
	fmt.Println("==========")
	ex3()
	fmt.Println("==========")
	ex4()
	fmt.Println("==========")
	ex5()
	fmt.Println("==========")
	ex6()
	fmt.Println("==========")
	ex7()
	fmt.Println("==========")
	ex8()
	fmt.Println("==========")
	ex9()
}
