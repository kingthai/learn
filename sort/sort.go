package sort

import (
	"fmt"
	"math/rand"
	"time"

	// "os"
	// "os/signal"
)

const (
	num      = 10
	rangeNum = 10
)

//func main() {
//	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
//	var buf []int
//	for i := 0; i < num; i++ {
//		buf = append(buf, randSeed.Intn(rangeNum))
//	}
//	t := time.Now()
//	//冒泡排序
//	// maopao(buf)
//	// 选择排序
//	// xuanze(buf)
//	// 插入排序
//	// charu(buf)
//	//希尔排序
//	// xier(buf)
//	//快速排序
//	// kuaisu(buf)
//	// 归并排序
//	// guibing(buf)
//	// 堆排序
//	duipai(buf)
//
//	// fmt.Println(buf)
//	fmt.Println(time.Since(t))
//
//	//等待退出
//	// c := make(chan os.Signal, 1)
//	// signal.Notify(c, os.Interrupt, os.Kill)
//	// <-c
//	// fmt.Println("Receive ctrl-c")
//}

var randSeed *rand.Rand
func  init ()  {
	randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 冒泡排序
func maopao(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 0; j < len(buf)-i-1; j++ {
			if buf[j] > buf[j+1] {
				times++
				buf[j], buf[j+1] = buf[j+1], buf[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
}

// 选择排序
func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
	}
	fmt.Println("xuanze times: ", times)
}

// 插入排序
func charu(buf []int) {
	times := 0
	for i := 1; i < len(buf); i++ {
		for j := i; j > 0; j-- {
			if buf[j] < buf[j-1] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
			} else {
				break
			}
		}
	}
	fmt.Println("charu times: ", times)
}

// 希尔排序
func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	// fmt.Println("buf: ", buf)
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					// fmt.Println("j: ", j, " data: ", buf[j], " j-incre: ", j-incre, " data: ", buf[j-incre])
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
				// fmt.Println("middle: ", buf)
			}
			// fmt.Println("outer: ", buf)
		}
		// fmt.Println("outer outer: ", buf, " incre: ", incre)

		if incre == 1 {
			break
		}
	}
	// fmt.Println("after: ", buf)
	fmt.Println("xier times: ", times)
}

// 快速排序
func kuaisu(buf []int) {
	quickSort(buf, 0, len(buf)-1)
}

// 2分法的最终版
func quickSort(a []int, l, r int) {
	if l < r {
		q := randomPartition(a, l, r)
		quickSort(a, q+1, r)
		quickSort(a, l, q-1)
	}
}

func randomPartition(a []int, l, r int) int {
	i := randSeed.Intn(r - l + 1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	key := a[r] // 选中放在数据的最后面的 标志元素
	i := l - 1
	for j:=l;j<r;j++ {
		if a[j] <= key { // 这里必须是<=
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i+1
}

//归并排序
func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
	// fmt.Printf("sort: buf: %v\n", a)
}

// 堆排序
func duipai(buf []int) {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}
