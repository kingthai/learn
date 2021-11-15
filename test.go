package main  
import (
	"fmt"
	"runtime"
	"sync"
)



func main() {
    //ch := make(chan int, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	runtime.GOMAXPROCS(1)

	go func() {
		defer wg.Done()
		for i:=1;i<20;i+=2 {
			fmt.Println(i)
			runtime.Gosched()
		}
	}()

	go func() {
		defer wg.Done()
    	for i:=0;i<20;i+=2 {
    		fmt.Println(i)
    		runtime.Gosched()
		}
	}()


	wg.Wait()

}


