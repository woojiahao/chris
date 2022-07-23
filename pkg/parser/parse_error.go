package parser

import (
	"fmt"
	"github.com/woojiahao/chris/pkg/lexer"
)

type errorReason string

const (
	invalidEndOfGroup            errorReason = "Expected ) at the end of a group, did not receive it"
	invalidEndOfFunction                     = "Expected ) at the end of a function call, did not receive it"
	invalidKeywordInFunctionCall             = "Function call must use keyword that has more than 1 character"
	invalidVariableInAssignment              = "Assignment variable must be a single-character value"
	invalidPrefixToken                       = "Invalid prefix token. Only valid prefix tokens are [<variable><keyword><number>-(]"
	invalidInfixToken                        = "Invalid infix token. Only valid infix tokens are [+-/*^(=]"
)

type ParseError struct {
	tokenType lexer.TokenType
	reason    errorReason
}

func (pe ParseError) Error() string {
	return fmt.Sprintf(
		"ParseError occurred on TokenType %s with reason %s",
		pe.tokenType.Name,
		pe.reason,
	)
}
