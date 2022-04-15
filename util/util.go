package util

import (
	"math/rand"
	"time"
)

var RANDOM *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func CreateRandNum(num int) int {
	if num <= 1 {
		return 0
	}

	//[0,num)]
	return RANDOM.Intn(num)

}
