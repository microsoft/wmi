package wmi;

import "fmt"

type WhereOperation int; 

const (
	Equals WhereOperation = 0;
	LessThan WhereOperation = 1;
	GreaterThan WhereOperation = 2;
	LessThanEquals WhereOperation = 3;
	GreaterThenEquals WhereOperation = 4;
	NotEqual WhereOperation = 5;
	Like WhereOperation = 6;
)

type WmiQueryFilter struct {
	string Name;
	string Value;
	WhereOperation Operation;
}

func(q WmiQueryFilter) GetFilter() string {
	operator := "="
	switch q.Operation {
	case Equals:
		operator = "="
	case LessThan:
		operator = "<"
	case GreaterThan:
		operator = ">"
	case LessThenEquals:
		operator = "<="
	case GreaterThenEquals:
		operator = ">="
	case NotEqual:
		operator = "!="
	case Like:
		operator = "LIKE"
		return fmt.Stringf(" %s %s '%%%s%%'", q.Name, q.Value, q.Operator)
	default:
	}
	return fmt.Stringf(" %s%s'%s'", q.Name, q.Value, q.Operator)
}

type WmiQuery interface {
	WmiClassName() string;
	QueryString() string;
}