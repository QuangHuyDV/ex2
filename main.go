package main

import (
	"context"
	"fmt"
	// "strconv"
	"time"
)

// 1. Tạo 1 chương trình cứ 3s in ra dòng chữ: time now: {milliseconds}: Sau khi in được 3 lần thì in ra dòng chữ kết thúc
func ex1() {
	timen := time.Now().UnixMilli()
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("time now", timen)

	}
	fmt.Println("kết thúc")
}

// 2. Việt 1 đoạn chương trình tính ra ngày hiện tại theo timestamp. Điểm mốc từ mức 0 tại 1/1/1970
func ex2() {
	time.Sleep(1*time.Second)
	timeq := time.Now().Unix()
	a := (((timeq / 60) / 60) / 24)
	fmt.Println("Day: ", a)
}

// 3. Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec Nhưng sau 3s thì kết thúc hàm đấy Sử dụng và tìm hiểu context. Nêu được tác dụng của context trong chương trình.
func ex3() {
	time.Sleep(1*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	// time.Sleep(3 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(3 * time.Second)
		}
	}
}

// 4. Tính ra số phút(number_of_minutes) của mốc thời gian sau 1592190294764144364
func ex4() {
	time.Sleep(1*time.Second)
	mi := 1592190294764144364 / 60000000000
	fmt.Println("So phut: ", int64(mi))
}

// 5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau 1592190385
func ex5() {
	time.Sleep(1*time.Second)
	a := 1592190385 / (60 * 60 * 24)
	day := time.Weekday(a)
	fmt.Printf("So ngay: %v, %v\n", a, day)

}

// 6. Trong golang mặc định thì thời gian dạng số được sử dụng với các loại mốc đơn vị nào?
func ex6() {
	time.Sleep(1*time.Second)
	fmt.Println("Đơn vị: second, nanosecond")
}

// 7. Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo ns chạy qua 1 hàm như sau. Sau 3s in ra hiệu của thời gian của thời gian hiện tại với biến dữ liệu trong context. in ra màn hình kết quả.

func ex7() {
	time.Sleep(1*time.Second)
	q := time.Now().UnixNano()
	ctx1 := context.WithValue(context.Background(), "ns", q)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	fmt.Println("Hiệu thời gian tính từ hiện tại: ")
	for {
		select {
		case <-ctx.Done():

			return
		default:
			time.Sleep(3 * time.Second)
			a2 := time.Now().UnixNano()
			va := ctx1.Value("ns")
			a1 := va.(int64)
			fmt.Println(a2 - a1)
		}
	}

}

// 8. Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ ${time.Now().Unix()} done
func ex8() {
	time.Sleep(1*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// time.Sleep(3 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end")
			return
		default:
			fmt.Printf("${%v}\n", time.Now().Unix())
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 9. Tạo 1 đoạn code sử dụng time.AfterFunc(), sau 100ms thì in ra dòng chữ i'm study
func ex9() {
	f := func() {
		fmt.Println("i'm study")
	}
	D := time.Duration(100) * time.Millisecond
	time1 := time.AfterFunc(D, f)
	defer time1.Stop()
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("ex1")
	ex1()
	fmt.Println("==========")
	fmt.Println("ex2")
	ex2()
	fmt.Println("==========")
	fmt.Println("ex3")
	ex3()
	fmt.Println("==========")
	fmt.Println("ex4")
	ex4()
	fmt.Println("==========")
	fmt.Println("ex5")
	ex5()
	fmt.Println("==========")
	fmt.Println("ex6")
	ex6()
	fmt.Println("==========")
	fmt.Println("ex7")
	ex7()
	fmt.Println("==========")
	fmt.Println("ex8")
	ex8()
	fmt.Println("==========")
	fmt.Println("ex9")
	ex9()
}
