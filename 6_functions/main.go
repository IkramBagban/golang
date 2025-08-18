package main

func sum2Num(num1, num2 int) int {
	return num1 + num2
}

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	// for i := 0; i < len(nums); i++ {
	// 	total = total + nums[i]
	// }
	return total
}

func sumAndSubstract(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func main() {
	println(sum2Num(6, 3))

	println(sumAll(2, 4, 6, 8))

	println(sumAndSubstract(9, 3))
}
