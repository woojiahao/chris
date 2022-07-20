package lexer

import (
	"strconv"
)

type Expression string

func (exp *Expression) Get(index int) *string {
	if index > len(*exp) {
		return nil
	}
	c := strconv.Itoa(int((*exp)[index]))
	return &c
}

func (exp *Expression) Len() int {
	return len(*exp)
}
