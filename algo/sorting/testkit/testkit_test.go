package testkit

import "testing"

func TestGetRandomKit(t *testing.T) {
	for i := 0; i < 100; i++ {
		testGetRandomKit(i, t)
	}
}

func BenchmarkGetRandomKit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomKit(10)
	}
}

func testGetRandomKit(size int, t *testing.T) {
	kit := GetRandomKit(size)

	if len(kit.N) != size {
		t.Error("1")
	}
	if len(kit.Want) != size {
		t.Error("2")
	}

	for i := 0; i < size-1; i++ {
		if kit.Want[i] > kit.Want[i+1] {
			t.Error("3")
		}
	}
}

func BenchmarkGetSortedKit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSortedKit(10)
	}
}
