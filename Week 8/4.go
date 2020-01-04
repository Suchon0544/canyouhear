package tot

func total(input []int) int {
	ans := []int{}
	sum := 0
	for i, v := range input {
		sum = sum + input
		ans = append(sum)
		if i == len(input) {
			break
		}
	}
	return sum
}
