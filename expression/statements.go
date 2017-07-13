package expression

// These strings match lexer tokens.
const (
	AssignStatementSelector = "%assign "
	CycleStatementSelector  = "{%cycle "
	LoopStatementSelector   = "%loop "
)

// A Statement is the result of parsing a string.
type Statement struct{ parseValue }

// Assignment returns a statement's assignment record.
func (s *Statement) Assignment() Assignment { return s.assgn }

// Expression returns a statement's expression.
func (s *Statement) Expression() Expression { return &expression{s.val} }

// Assignment captures the parse of an {% assign %} statement
type Assignment struct {
	Name    string
	ValueFn Expression
}

// Loop captures the parse of a {% loop %} statement
type Loop struct {
	Variable string
	Expr     interface{}
	loopModifiers
}

type loopModifiers struct {
	Limit    *int
	Offset   int
	Reversed bool
}

// ParseStatement parses an statement into an Expression that can evaluated to return a
// structure specific to the statement.
func ParseStatement(sel, source string) (*Statement, error) {
	p, err := parse(sel + source)
	if err != nil {
		return nil, err
	}
	return &Statement{*p}, nil
}
