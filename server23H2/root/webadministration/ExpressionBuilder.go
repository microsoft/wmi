// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ExpressionBuilder struct
type ExpressionBuilder struct {
	*CollectionElement

	//
	ExpressionPrefix string

	//
	Type string
}

func NewExpressionBuilderEx1(instance *cim.WmiInstance) (newInstance *ExpressionBuilder, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ExpressionBuilder{
		CollectionElement: tmp,
	}
	return
}

func NewExpressionBuilderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ExpressionBuilder, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ExpressionBuilder{
		CollectionElement: tmp,
	}
	return
}

// SetExpressionPrefix sets the value of ExpressionPrefix for the instance
func (instance *ExpressionBuilder) SetPropertyExpressionPrefix(value string) (err error) {
	return instance.SetProperty("ExpressionPrefix", (value))
}

// GetExpressionPrefix gets the value of ExpressionPrefix for the instance
func (instance *ExpressionBuilder) GetPropertyExpressionPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("ExpressionPrefix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *ExpressionBuilder) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *ExpressionBuilder) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
