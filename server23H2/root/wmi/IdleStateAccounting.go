// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// IdleStateAccounting struct
type IdleStateAccounting struct {
	*cim.WmiInstance

	//
	FailedTransitions uint32

	//
	IdleTimeBuckets []uint32

	//
	IdleTransitions uint32

	//
	InvalidBucketIndex uint32

	//
	TotalTime uint64
}

func NewIdleStateAccountingEx1(instance *cim.WmiInstance) (newInstance *IdleStateAccounting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &IdleStateAccounting{
		WmiInstance: tmp,
	}
	return
}

func NewIdleStateAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IdleStateAccounting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IdleStateAccounting{
		WmiInstance: tmp,
	}
	return
}

// SetFailedTransitions sets the value of FailedTransitions for the instance
func (instance *IdleStateAccounting) SetPropertyFailedTransitions(value uint32) (err error) {
	return instance.SetProperty("FailedTransitions", (value))
}

// GetFailedTransitions gets the value of FailedTransitions for the instance
func (instance *IdleStateAccounting) GetPropertyFailedTransitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedTransitions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIdleTimeBuckets sets the value of IdleTimeBuckets for the instance
func (instance *IdleStateAccounting) SetPropertyIdleTimeBuckets(value []uint32) (err error) {
	return instance.SetProperty("IdleTimeBuckets", (value))
}

// GetIdleTimeBuckets gets the value of IdleTimeBuckets for the instance
func (instance *IdleStateAccounting) GetPropertyIdleTimeBuckets() (value []uint32, err error) {
	retValue, err := instance.GetProperty("IdleTimeBuckets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetIdleTransitions sets the value of IdleTransitions for the instance
func (instance *IdleStateAccounting) SetPropertyIdleTransitions(value uint32) (err error) {
	return instance.SetProperty("IdleTransitions", (value))
}

// GetIdleTransitions gets the value of IdleTransitions for the instance
func (instance *IdleStateAccounting) GetPropertyIdleTransitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleTransitions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInvalidBucketIndex sets the value of InvalidBucketIndex for the instance
func (instance *IdleStateAccounting) SetPropertyInvalidBucketIndex(value uint32) (err error) {
	return instance.SetProperty("InvalidBucketIndex", (value))
}

// GetInvalidBucketIndex gets the value of InvalidBucketIndex for the instance
func (instance *IdleStateAccounting) GetPropertyInvalidBucketIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InvalidBucketIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTotalTime sets the value of TotalTime for the instance
func (instance *IdleStateAccounting) SetPropertyTotalTime(value uint64) (err error) {
	return instance.SetProperty("TotalTime", (value))
}

// GetTotalTime gets the value of TotalTime for the instance
func (instance *IdleStateAccounting) GetPropertyTotalTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
