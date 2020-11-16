package iteration

// Repeat will return a string repeated the number of times requested by count
func Repeat(str string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += str
	}
	return repeated
}
