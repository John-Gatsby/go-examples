package sorting

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

type kit struct {
	random Array
	sorted Array
}

var once sync.Once

func getKit() kit {
	kit := kit{}

	once.Do(func() {
		rand.Seed(time.Now().Unix())
	})

	for i := 0; i < arrSize; i++ {
		n := rand.Intn(maxNumber)
		kit.random[i] = n
		kit.sorted[i] = n
	}

	sort.Ints(kit.sorted[:])

	return kit
}

func TestBubble(t *testing.T) {
	kit := getKit()
	sorted := Bubble(kit.random)
	if sorted != kit.sorted {
		t.Errorf("error sorting Bubble:\n should be: %v\n got: %v", kit.sorted, sorted)
	} else {
		t.Logf("Bubble:\n should be: %v\n got: %v", kit.sorted, sorted)
	}
}

func BenchmarkBubble(b *testing.B) {
	kit := getKit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bubble(kit.random)
	}
}
