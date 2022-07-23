package parser

import (
	"fmt"
	"github.com/woojiahao/chris/pkg/lexer"
)

type ParseError struct {
	tokenType lexer.TokenType
	reason    string
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf(
		"ParseError occurred on TokenType %s with reason %s",
		pe.tokenType.Name,
		pe.reason,
	)
}
