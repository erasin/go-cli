package main

import (
	"fmt"
)

func main() {
	bubbles := []int{15, 23, 733, 47, 85, 675, 678, 8, 988, 10, 11, 8126, 173}
	maopao1(bubbles)
	maopao2(bubbles)

}

func maopao1(bubbles []int) {
	leng := len(bubbles) - 1 //数组长度
	for i := 0; i < leng; i++ {
		for i2 := 0; i2 < leng-i; i2++ {
			if bubbles[i2] > bubbles[i2+1] {
				//数值交换
				bubbles[i2] = bubbles[i2] + bubbles[i2+1]
				bubbles[i2+1] = bubbles[i2] - bubbles[i2+1]
				bubbles[i2] = bubbles[i2] - bubbles[i2+1]
			}
		}

	}
	fmt.Println(bubbles)
}

//对于没有发生过排序交换的数组，可判定为已按顺序排序，结束排序返回结果。
//如果后面元素没有发生过排序，说明后面部分已经按顺序排列，可以记录最后进行交换的位置，不对此位置后的元素再进行排序。

func maopao2(bubbles []int) {
	leng := len(bubbles) - 1 //数组长度
	lIndex := leng           //最后交换位置
	swap := false            //是否交换过
	for i := 0; i < leng; i++ {
		tmpIndex := 0
		for i2 := 0; i2 < lIndex; i2++ {
			if bubbles[i2] > bubbles[i2+1] {
				//数值交换
				bubbles[i2] = bubbles[i2] + bubbles[i2+1]
				bubbles[i2+1] = bubbles[i2] - bubbles[i2+1]
				bubbles[i2] = bubbles[i2] - bubbles[i2+1]
				tmpIndex = i2
				swap = true
			}
		}
		lIndex = tmpIndex
		if !swap {
			break
		}

	}
	fmt.Println(bubbles)
}
