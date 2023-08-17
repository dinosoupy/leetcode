func reverse(x int) int {
    var rev int
	for x != 0 {
		a := x % 10
		x = x / 10
		rev = rev*10 + a
	}
    if rev < -2147483648 || rev > 2147483648 {
    return 0
    }
	return rev
}