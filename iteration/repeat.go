package iteration

import "strings"

const defaultRepeatCount = 5

func Repeat(character string, timesArg ...int) string {
	var repeated strings.Builder

	repeatCount := defaultRepeatCount
	if len(timesArg) > 0 {
		repeatCount = timesArg[0]
	}

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
