package iteration

func Repeat(character string, numOfIterations int) string {
	var repeated string
	for i := 0; i < numOfIterations; i++ {
		repeated += character
	}
	return repeated
}
