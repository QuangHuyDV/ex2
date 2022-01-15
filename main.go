package main

import (
	"context"
	"fmt"
	"time"
)

func ex1() {
	timen := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("time now", timen)
		if i == 3 {
			fmt.Println("kết thúp")
		}
	}
}

func ex2() {
	timeq := time.Now().Unix()
	a := (((timeq / 60) / 60) / 24) / 7
	fmt.Println("Số tuần: ", a)
}

func text(ctx context.Context) {
	fmt.Println("text")
	time.Sleep(1 * time.Second)
}

func ex3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("sleep: ", i)
		text(ctx)
	}
}

func ex4() {
	mi := 1592190294764144364 / 60000000000
	fmt.Println("So phut: ", int64(mi))
	fmt.Println(time.Unix(int64(mi), 0))
}

func ex5() {
	a := 1592190385 / (60 * 60 * 24)
	fmt.Println("So ngay: ", a)

}

func ex6() {
	fmt.Println("Đơn vị: second, nanosecond")
}

func ex7() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	x(ctx)

}

func x(ctx context.Context) {
	time1 := time.Now().UnixNano()
	fmt.Println("Thời gian hiện tại:", time1)
	time.Sleep(3 * time.Second)
	time2 := time.Now().UnixNano()
	fmt.Println("Thời gian sau 3s: ", time2)
	hieu := time2 - time1
	println("Hiệu thời gian sau 3s:", hieu)
}

func test(a, quit chan int64) {
	time11 := 1
	y := "done"
	for {
		select {
		case a <- int64(time11):
			fmt.Printf("")
		case <-quit:
			fmt.Println(y)
			return
		}
	}
}

func ex8() {
	a := make(chan int64)
	quit := make(chan int64)
	go func() {
		for i := 0; i < 1; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("${time.Now().Unix()}")
		}
		quit <- 0
	}()
	test(a, quit)

}

func ex9() {
	f := func() {
		fmt.Println("i'm study")
	}
	D := time.Duration(100) * time.Millisecond
	time1 := time.AfterFunc(D, f)
	defer time1.Stop()
	time.Sleep(2*time.Second)
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
