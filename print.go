package main

import (
	"fmt"
	"runtime"
	"sync"
)

var zf []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func PrintZF() {
	defer wg.Done()
	for i := 0; i < 26; i++ {
		fmt.Println(zf[i])
		runtime.Gosched()
	}

}
func PrintN() {
	defer wg.Done()
	for i := 0; i < 26; i++ {
		fmt.Println(i)
		runtime.Gosched()
	}

}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(1)
	go PrintZF()
	wg.Add(1)
	go PrintN()
	wg.Wait()
}