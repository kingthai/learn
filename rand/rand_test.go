package rand

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	a := []int{1,2,3,4,5}
	Shuffle(a)
	fmt.Println(a)
}