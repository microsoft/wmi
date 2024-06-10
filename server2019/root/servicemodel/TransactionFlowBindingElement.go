// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TransactionFlowBindingElement struct
type TransactionFlowBindingElement struct {
	*BindingElement

	// Indicates whether support for wildcard action with transaction flow is enabled.
	AllowWildcardAction bool

	// Indicates whether incoming transaction flow is enabled.
	TransactionFlow bool

	// The transaction protocol used in flowing a transaction.
	TransactionProtocol string
}

func NewTransactionFlowBindingElementEx1(instance *cim.WmiInstance) (newInstance *TransactionFlowBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransactionFlowBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewTransactionFlowBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransactionFlowBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransactionFlowBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetAllowWildcardAction sets the value of AllowWildcardAction for the instance
func (instance *TransactionFlowBindingElement) SetPropertyAllowWildcardAction(value bool) (err error) {
	return instance.SetProperty("AllowWildcardAction", (value))
}

// GetAllowWildcardAction gets the value of AllowWildcardAction for the instance
func (instance *TransactionFlowBindingElement) GetPropertyAllowWildcardAction() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowWildcardAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTransactionFlow sets the value of TransactionFlow for the instance
func (instance *TransactionFlowBindingElement) SetPropertyTransactionFlow(value bool) (err error) {
	return instance.SetProperty("TransactionFlow", (value))
}

// GetTransactionFlow gets the value of TransactionFlow for the instance
func (instance *TransactionFlowBindingElement) GetPropertyTransactionFlow() (value bool, err error) {
	retValue, err := instance.GetProperty("TransactionFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTransactionProtocol sets the value of TransactionProtocol for the instance
func (instance *TransactionFlowBindingElement) SetPropertyTransactionProtocol(value string) (err error) {
	return instance.SetProperty("TransactionProtocol", (value))
}

// GetTransactionProtocol gets the value of TransactionProtocol for the instance
func (instance *TransactionFlowBindingElement) GetPropertyTransactionProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionProtocol")
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
