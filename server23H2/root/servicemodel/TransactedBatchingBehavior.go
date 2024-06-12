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

// TransactedBatchingBehavior struct
type TransactedBatchingBehavior struct {
	*Behavior

	// The maximum batch size for the transacted batching.
	MaxBatchSize int32
}

func NewTransactedBatchingBehaviorEx1(instance *cim.WmiInstance) (newInstance *TransactedBatchingBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TransactedBatchingBehavior{
		Behavior: tmp,
	}
	return
}

func NewTransactedBatchingBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TransactedBatchingBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TransactedBatchingBehavior{
		Behavior: tmp,
	}
	return
}

// SetMaxBatchSize sets the value of MaxBatchSize for the instance
func (instance *TransactedBatchingBehavior) SetPropertyMaxBatchSize(value int32) (err error) {
	return instance.SetProperty("MaxBatchSize", (value))
}

// GetMaxBatchSize gets the value of MaxBatchSize for the instance
func (instance *TransactedBatchingBehavior) GetPropertyMaxBatchSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBatchSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
