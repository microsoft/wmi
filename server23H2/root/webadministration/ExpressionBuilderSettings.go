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

// ExpressionBuilderSettings struct
type ExpressionBuilderSettings struct {
	*EmbeddedObject

	//
	ExpressionBuilders []ExpressionBuilder
}

func NewExpressionBuilderSettingsEx1(instance *cim.WmiInstance) (newInstance *ExpressionBuilderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ExpressionBuilderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewExpressionBuilderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ExpressionBuilderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ExpressionBuilderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetExpressionBuilders sets the value of ExpressionBuilders for the instance
func (instance *ExpressionBuilderSettings) SetPropertyExpressionBuilders(value []ExpressionBuilder) (err error) {
	return instance.SetProperty("ExpressionBuilders", (value))
}

// GetExpressionBuilders gets the value of ExpressionBuilders for the instance
func (instance *ExpressionBuilderSettings) GetPropertyExpressionBuilders() (value []ExpressionBuilder, err error) {
	retValue, err := instance.GetProperty("ExpressionBuilders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ExpressionBuilder)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ExpressionBuilder is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ExpressionBuilder(valuetmp))
	}

	return
}
