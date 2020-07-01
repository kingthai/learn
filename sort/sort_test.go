package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMaopao(t *testing.T) {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	start := time.Now()
	maopao(buf)
	fmt.Println(time.Since(start), buf)

}
