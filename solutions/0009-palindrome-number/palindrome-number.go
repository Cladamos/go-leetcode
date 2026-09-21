package solutions

func isPalindrome(x int) bool {
	num := x
	var reversed int
	for num > 0 {
		mod := num % 10
		num /= 10
		reversed *= 10
		reversed += mod
	}
	return reversed == x
}
