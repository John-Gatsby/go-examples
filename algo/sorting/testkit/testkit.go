package testkit

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

const maxNumber = 100

var once sync.Once

type Kit struct {
	N    []int
	Want []int
}

func GetRandomKit(size int) Kit {
	kit := Kit{
		N:    make([]int, size),
		Want: make([]int, size),
	}

	once.Do(func() {
		rand.Seed(time.Now().Unix())
	})

	for i := 0; i < size; i++ {
		n := rand.Intn(maxNumber)
		kit.N[i] = n
		kit.Want[i] = n
	}

	sort.Ints(kit.Want)

	return kit
}

func GetSortedKit(size int) Kit {
	kit := GetRandomKit(size)
	copy(kit.N, kit.Want)
	return kit
}
