package main

func ReverseNumber(num int) int {
	ans := 0
	for num != 0 {
		ans = ans*10 + num%10
		num /= 10
	}
	return ans
}

func ReverseData(arr [5]int) [5]int {
	for i := 0; i < 5; i++ {
		arr[i] = ReverseNumber(arr[i])
	}
	arr[0], arr[4] = arr[4], arr[0]
	arr[1], arr[3] = arr[3], arr[1]
	return arr // TODO: replace this
}

func main() {
	data := [5]int{12, 23, 34, 45, 56}
	ReverseData(data)
}
