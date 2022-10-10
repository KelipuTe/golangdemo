package orm_select

// operator 代表操作符
type operator string

func (this operator) String() string {
	return string(this)
}

const (
	opEQ  operator = "="
	opGT  operator = ">"
	opLT  operator = "<"
	opAND operator = "AND"
	opOR  operator = "OR"
	opNOT operator = "NOT"
)
