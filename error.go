package simpleregexengine

import "fmt"

type ParseErrorCode string

var SyntaxError ParseErrorCode = "SyntaxError"

const CompilationError = "CompilationError"

type RegexError struct {
	Code    ParseErrorCode
	Message string
	Pos     int
}

func (p *RegexError) Error() string {
	return fmt.Sprintf("code=%s, message=%s, pos=%d", p.Code, p.Message, p.Pos)
}
