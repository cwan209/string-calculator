package stringCalculator

import (
	"fmt"
	"regexp"
	"strconv"
)

func Add(s string) int {
	if len(s) > 1 {
		var patternString string
		if s[:2] == "//" {
			newD := s[2]			patternString = fmt.Sprintf(`[,%c\n]`, newD)
			fmt.Println(patternString)
			fmt.Println(patternString == `[,;\n]`)
		} else {
			patternString = `[,\n]`
		}
		pattern := regexp.MustCompile(patternString)
		adders := pattern.Split(s, -1)
		var sum int
		for _, adder := range adders {
			adder2, _ := strconv.Atoi(adder)
			sum += adder2
		}
		return sum
	}

	i, _ := strconv.Atoi(s)
	return i
}
