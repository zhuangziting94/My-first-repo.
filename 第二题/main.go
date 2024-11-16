package main

import "fmt"

func main() {
	var n int
	fmt.Printf("请输入数组的长度：%d", n)
	var sl []int
	fmt.Printf("请输入数组：")
	fmt.Scan(sl)
	n = len(sl)
	var sum int
	var Sum int
	for i := 0; i < n; i++ {
		switch {
		case sl[i] > sl[i+1]:
			sum = twoNumber(sl[i], sl[i+1])
		case sl[i] == sl[i+1]:
			sum = 0
		case sl[i] < sl[i+1]:
			sum = sl[i+1] - sl[i]
		}
		Sum += sum
	}
	fmt.Println(Sum)
}

func twoNumber(t1 int, t2 int) int {
	var sum1 int
	if t1-t2-2 >= 0 {
		sum1 = t1 - t2 - 2

	} else {
		sum1 = 2 + t2 - t1
	}
	return sum1
}
