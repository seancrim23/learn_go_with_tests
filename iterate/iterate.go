package iterate

func Repeat(letter string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}
	return repeated
}
