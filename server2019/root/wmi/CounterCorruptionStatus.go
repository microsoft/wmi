// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CounterCorruptionStatus struct
type CounterCorruptionStatus struct {
	*cim.WmiInstance

	//
	LastKnownGoodTimestamp uint64

	//
	ProfileSource uint32
}

func NewCounterCorruptionStatusEx1(instance *cim.WmiInstance) (newInstance *CounterCorruptionStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CounterCorruptionStatus{
		WmiInstance: tmp,
	}
	return
}

func NewCounterCorruptionStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CounterCorruptionStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CounterCorruptionStatus{
		WmiInstance: tmp,
	}
	return
}

// SetLastKnownGoodTimestamp sets the value of LastKnownGoodTimestamp for the instance
func (instance *CounterCorruptionStatus) SetPropertyLastKnownGoodTimestamp(value uint64) (err error) {
	return instance.SetProperty("LastKnownGoodTimestamp", (value))
}

// GetLastKnownGoodTimestamp gets the value of LastKnownGoodTimestamp for the instance
func (instance *CounterCorruptionStatus) GetPropertyLastKnownGoodTimestamp() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastKnownGoodTimestamp")
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

// SetProfileSource sets the value of ProfileSource for the instance
func (instance *CounterCorruptionStatus) SetPropertyProfileSource(value uint32) (err error) {
	return instance.SetProperty("ProfileSource", (value))
}

// GetProfileSource gets the value of ProfileSource for the instance
func (instance *CounterCorruptionStatus) GetPropertyProfileSource() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProfileSource")
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
