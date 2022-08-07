package sorting

type Ordered interface {
	~int | ~string | ~int32 | ~int64
}

func BubbleSort[K Ordered](numbers []K) []K {
	var temp K
	length := len(numbers)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {

			if numbers[j] > numbers[j+1] {
				temp = numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	return numbers
}
