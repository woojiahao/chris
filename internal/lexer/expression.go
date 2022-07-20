package lexer

import (
	"strconv"
	"unicode"
)

type Expression string

func (exp *Expression) Get(index int) *rune {
	if index > len(*exp) {
		return nil
	}
	c := []rune(*exp)[index]
	return &c
}

// substring retrieves the substring from start to end (both inclusive)
func (exp *Expression) substring(start, end int) *string {
	if end >= len(*exp) {
		end = len(*exp) - 1
	}

	if start < 0 {
		start = 0
	}

	substring := (*exp)[start : end+1]
	casted := string(substring)

	return &casted
}

func (exp *Expression) readNumber(start int) (*float64, int) {
	i := start
	for {
		cur := exp.Get(start)

		if cur == nil {
			return nil, -1
		}

		if unicode.IsDigit(*cur) {
			i++
		} else {
			break
		}
	}

	// TODO: Add unit test for this
	numStr := *exp.substring(start, i)
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return nil, -1
	}

	return &num, i + 1
}

func (exp *Expression) Len() int {
	return len(*exp)
}
