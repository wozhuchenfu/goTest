package suanfa

//最大不重复子串长度计算
func LengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,ch := range []byte(s){
		if lastI,ok := lastOccurred[ch];ok&&lastI>=start {
			start = lastI + 1
		}
		if i-start+1>maxLength {
			maxLength = i-start+1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
