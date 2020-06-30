package rand

import (
	"math/rand"
	"time"
)

func Shuffle(items[] int) {
	rand.Seed(time.Now().UnixNano())
	i := len(items) - 1
	for i >= 1 {
		j := rand.Int63n(int64(i + 1))
		items[i], items[j] = items[j], items[i]
		i--
	}

}