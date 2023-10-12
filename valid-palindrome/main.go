func isPalindrome(s string) bool {
  s = strings.ToLower(s)
	sbyte := []byte(s)
	start, end := 0, len(s)-1

	for start < end {
		for start < end && !isAlphaNumeric(sbyte[start]) {
			start++
		}
		for start < end && !isAlphaNumeric(sbyte[end]) {
			end--
		}
		if sbyte[start] != sbyte[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphaNumeric(char byte) bool {
	return ('a' <= char && char <= 'z') ||
		('A' <= char && char <= 'Z') ||
		('0' <= char && char <= '9')
}