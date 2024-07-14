package arrays

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = []int{}

	for _, subArr := range numbersToSum {
		sums = append(sums, Sum(subArr))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sums = []int{}

	for _, subArr := range numbersToSum {
		if len(subArr) > 0 {
			tail := subArr[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return
}
