package lexer

type Expression string

func (exp *Expression) Get(index int) *rune {
	if index > len(*exp) {
		return nil
	}
	c := []rune(*exp)[index]
	return &c
}

func (exp *Expression) Len() int {
	return len(*exp)
}
