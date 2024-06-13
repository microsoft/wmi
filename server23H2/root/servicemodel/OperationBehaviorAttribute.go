// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// OperationBehaviorAttribute struct
type OperationBehaviorAttribute struct {
	*Behavior

	// The state of the auto dispose feature for parameters.
	AutoDisposeParameters bool

	// Indicates the level of caller impersonation that the operation supports.
	Impersonation string

	// Indicates when in the course of an operation invocation to recycle the object.
	ReleaseInstanceMode string

	// Indicates whether to automatically commit the current transaction if no unhandled exceptions occur.
	TransactionAutoComplete bool

	// Indicates whether the operation requires a transaction.
	TransactionScopeRequired bool
}

func NewOperationBehaviorAttributeEx1(instance *cim.WmiInstance) (newInstance *OperationBehaviorAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OperationBehaviorAttribute{
		Behavior: tmp,
	}
	return
}

func NewOperationBehaviorAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OperationBehaviorAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OperationBehaviorAttribute{
		Behavior: tmp,
	}
	return
}

// SetAutoDisposeParameters sets the value of AutoDisposeParameters for the instance
func (instance *OperationBehaviorAttribute) SetPropertyAutoDisposeParameters(value bool) (err error) {
	return instance.SetProperty("AutoDisposeParameters", (value))
}

// GetAutoDisposeParameters gets the value of AutoDisposeParameters for the instance
func (instance *OperationBehaviorAttribute) GetPropertyAutoDisposeParameters() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoDisposeParameters")
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

// SetImpersonation sets the value of Impersonation for the instance
func (instance *OperationBehaviorAttribute) SetPropertyImpersonation(value string) (err error) {
	return instance.SetProperty("Impersonation", (value))
}

// GetImpersonation gets the value of Impersonation for the instance
func (instance *OperationBehaviorAttribute) GetPropertyImpersonation() (value string, err error) {
	retValue, err := instance.GetProperty("Impersonation")
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

// SetReleaseInstanceMode sets the value of ReleaseInstanceMode for the instance
func (instance *OperationBehaviorAttribute) SetPropertyReleaseInstanceMode(value string) (err error) {
	return instance.SetProperty("ReleaseInstanceMode", (value))
}

// GetReleaseInstanceMode gets the value of ReleaseInstanceMode for the instance
func (instance *OperationBehaviorAttribute) GetPropertyReleaseInstanceMode() (value string, err error) {
	retValue, err := instance.GetProperty("ReleaseInstanceMode")
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

// SetTransactionAutoComplete sets the value of TransactionAutoComplete for the instance
func (instance *OperationBehaviorAttribute) SetPropertyTransactionAutoComplete(value bool) (err error) {
	return instance.SetProperty("TransactionAutoComplete", (value))
}

// GetTransactionAutoComplete gets the value of TransactionAutoComplete for the instance
func (instance *OperationBehaviorAttribute) GetPropertyTransactionAutoComplete() (value bool, err error) {
	retValue, err := instance.GetProperty("TransactionAutoComplete")
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

// SetTransactionScopeRequired sets the value of TransactionScopeRequired for the instance
func (instance *OperationBehaviorAttribute) SetPropertyTransactionScopeRequired(value bool) (err error) {
	return instance.SetProperty("TransactionScopeRequired", (value))
}

// GetTransactionScopeRequired gets the value of TransactionScopeRequired for the instance
func (instance *OperationBehaviorAttribute) GetPropertyTransactionScopeRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("TransactionScopeRequired")
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
