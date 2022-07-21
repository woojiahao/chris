package lexer

import (
	"strconv"
	"unicode"
)

type Expression string

func (exp Expression) Get(index int) *rune {
	if index >= exp.Len() {
		return nil
	}
	c := []rune(exp)[index]
	return &c
}

// substring retrieves the substring from start to end (both inclusive)
func (exp Expression) substring(start, end int) string {
	if end >= exp.Len() {
		end = exp.Len() - 1
	}

	if start < 0 {
		start = 0
	}

	substring := exp[start : end+1]
	casted := string(substring)

	return casted
}

func (exp Expression) readNumber(start int) (*float64, int) {
	i := start
	for {
		cur := exp.Get(i)

		if cur == nil && i < exp.Len() {
			// If the cur is nil and we're not at the end yet, return
			return nil, -1
		} else if cur == nil && i >= exp.Len() {
			break
		}

		if unicode.IsDigit(*cur) {
			i++
		} else {
			break
		}
	}

	// TODO: Add unit test for this
	numStr := exp.substring(start, i-1)
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return nil, -1
	}

	return &num, i
}

func (exp Expression) Len() int {
	return len(exp)
}
