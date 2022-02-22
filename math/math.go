package math

func Sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func Multiply(i ...int) int {
	res := 1
	for _, v := range i {
		res *= v
	}
	return res
}
