// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TransactionFlowAttribute struct
type TransactionFlowAttribute struct {
	*Behavior

	// Indicates whether transactions flow or not.
	TransactionFlowOption string
}

func NewTransactionFlowAttributeEx1(instance *cim.WmiInstance) (newInstance *TransactionFlowAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransactionFlowAttribute{
		Behavior: tmp,
	}
	return
}

func NewTransactionFlowAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransactionFlowAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransactionFlowAttribute{
		Behavior: tmp,
	}
	return
}

// SetTransactionFlowOption sets the value of TransactionFlowOption for the instance
func (instance *TransactionFlowAttribute) SetPropertyTransactionFlowOption(value string) (err error) {
	return instance.SetProperty("TransactionFlowOption", (value))
}

// GetTransactionFlowOption gets the value of TransactionFlowOption for the instance
func (instance *TransactionFlowAttribute) GetPropertyTransactionFlowOption() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionFlowOption")
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
