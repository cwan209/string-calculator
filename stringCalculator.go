package stringCalculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Add(s string) (int, error) {
	if len(s) > 1 {
		var patternString string
		if s[:2] == "//" {
			newD := s[2]
			patternString = fmt.Sprintf(`[,%c\n]`, newD)
		} else {
			patternString = `[,\n]`
		}
		pattern := regexp.MustCompile(patternString)
		adders := pattern.Split(s, -1)
		var sum int
		for _, adder := range adders {
			adder2, _ := strconv.Atoi(adder)
			if adder2 < 0 {
				return 0, errors.New("negatives not allowed")
			}
			if adder2 >= 1000 {
				continue
			}
			sum += adder2
		}
		return sum, nil
	}

	i, _ := strconv.Atoi(s)
	return i, nil
}
