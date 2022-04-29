package testkit

import (
	"math/rand"
	"sort"
	"time"
)

const maxNumber = 100

type Kit struct {
	N    []int
	Want []int
}

func GetRandomKit(size int) Kit {
	kit := Kit{
		N:    make([]int, size),
		Want: make([]int, size),
	}

	rand.Seed(time.Now().Unix())

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
