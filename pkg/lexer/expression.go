package lexer

import (
	"github.com/woojiahao/chris/internal/utils"
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
	i, decimalRead := start, false
	for i < exp.Len() {
		if cur := exp.Get(i); cur == nil {
			return nil, -1
		} else if unicode.IsDigit(*cur) {
			i++
		} else if *cur == '.' {
			if decimalRead {
				return nil, -1
			} else {
				decimalRead = true
				i++
			}
		} else {
			break
		}
	}

	if num := utils.StrToFloat(exp.substring(start, i-1)); num == nil {
		return nil, -1
	} else {
		return num, i
	}
}

func (exp Expression) readWord(start int) (*string, int) {
	i := exp.read(start, unicode.IsLetter)
	if i == -1 {
		return nil, -1
	}

	word := exp.substring(start, i-1)
	return &word, i
}

func (exp Expression) Len() int {
	return len(exp)
}

// read continues to read the next character until a given predicate is false
func (exp Expression) read(i int, predicate func(rune) bool) int {
	for i < exp.Len() {
		cur := exp.Get(i)
		if cur == nil {
			return -1
		}

		if predicate(*cur) {
			i++
		} else {
			break
		}
	}

	return i
}
