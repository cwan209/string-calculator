package stringCalculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Add(s string) (int, error) {
	if len(s) > 1 {
		var patternString string
		if s[:3] == "//[" {
			// loop through
			var delimiters []string
			start := false
			curr := ""
			for i := 2; i < len(s); i++ {
				c := s[i]
				if c == '[' {
					start = true
				} else if c == ']' {
					start = false
					delimiters = append(delimiters, curr)
					curr = ""
				} else if start {
					curr = curr + string(c)
				}
			}

			patternString = "[,"
			for _, delimiter := range delimiters {
				patternString += "|"
				patternString += delimiter
			}
			patternString += "|\n]"
			s = strings.TrimPrefix(s, "//")
			s = regexp.MustCompile(`^\[.*?\]`).ReplaceAllString(s, "")
		} else if s[:2] == "//" {
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
