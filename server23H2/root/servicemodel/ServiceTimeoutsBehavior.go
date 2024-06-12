// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceTimeoutsBehavior struct
type ServiceTimeoutsBehavior struct {
	*Behavior

	// The period within which a transaction must complete.
	TransactionTimeout string
}

func NewServiceTimeoutsBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceTimeoutsBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceTimeoutsBehavior{
		Behavior: tmp,
	}
	return
}

func NewServiceTimeoutsBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceTimeoutsBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceTimeoutsBehavior{
		Behavior: tmp,
	}
	return
}

// SetTransactionTimeout sets the value of TransactionTimeout for the instance
func (instance *ServiceTimeoutsBehavior) SetPropertyTransactionTimeout(value string) (err error) {
	return instance.SetProperty("TransactionTimeout", (value))
}

// GetTransactionTimeout gets the value of TransactionTimeout for the instance
func (instance *ServiceTimeoutsBehavior) GetPropertyTransactionTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionTimeout")
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
