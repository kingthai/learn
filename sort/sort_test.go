package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMaopao(t *testing.T) {
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	fmt.Println("源数组: ", buf)
	start := time.Now()
	//maopao(buf)
	//kuaisu(buf)
	//a := guibing(buf)
	duipai(buf)
	fmt.Println(time.Since(start))
	fmt.Println("结果数组:" , buf)

}
