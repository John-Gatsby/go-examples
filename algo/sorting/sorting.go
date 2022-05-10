package sorting

const (
	arrSize   = 10
	maxNumber = 100
)

type Array = [arrSize]int

func Bubble(arr Array) Array {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
