package insertionSort

import (
	"math/rand"
	"time"
)

func GenerateOrderedArray(n int, bound int) []int {
	arr := make([]int, 0, 0)
	//加入随机数生产数组
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		//生成一个长度为n（i）的数组，数字范围在[0,bound)的数组
		arr = append(arr, rand.Intn(bound))
	}
	return arr
}
