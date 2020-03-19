// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package query

import (
	"fmt"
)

type CompareOperator string

const (
	Equals            CompareOperator = "="
	LessThan          CompareOperator = "<"
	GreaterThan       CompareOperator = ">"
	LessThanEquals    CompareOperator = "<="
	GreaterThanEquals CompareOperator = ">="
	NotEquals         CompareOperator = "<>"
	Like              CompareOperator = "LIKE"
)

type WmiQueryFilter struct {
	Name     string
	Value    string
	Operator CompareOperator
}

type WmiQuery struct {
	ClassName string
	Filters   []*WmiQueryFilter
}

func NewWmiQuery(className string, filters ...string) (wquery *WmiQuery) {
	wquery = &WmiQuery{ClassName: className, Filters: []*WmiQueryFilter{}}
	if len(filters) == 0 {
		return
	}

	for i := 0; i < len(filters); i = i + 2 {
		qfilter := NewWmiQueryFilter(filters[i], filters[i+1], Equals)
		wquery.Filters = append(wquery.Filters, qfilter)
	}

	return
}

// NewWmiQueryFilter
func NewWmiQueryFilter(name, value string, oper CompareOperator) *WmiQueryFilter {
	return &WmiQueryFilter{Name: name, Value: value, Operator: oper}
}

func (q *WmiQueryFilter) String() string {
	if q.Operator == Like {
		return fmt.Sprintf("%s %s '%%%s%%'", q.Name, q.Operator, q.Value)
	} else {
		return fmt.Sprintf("%s %s '%s'", q.Name, q.Operator, q.Value)
	}
}

func (q *WmiQuery) String() (queryString string) {
	queryString = fmt.Sprintf("SELECT * FROM %s", q.ClassName)

	if len(q.Filters) == 0 {
		return
	}

	queryString = fmt.Sprintf("%s WHERE ", queryString)

	for _, val := range q.Filters[:len(q.Filters)-1] {
		queryString = queryString + fmt.Sprintf(" %s AND", val.String())
	}

	queryString = queryString + fmt.Sprintf(" %s ", q.Filters[len(q.Filters)-1].String())
	return
}
