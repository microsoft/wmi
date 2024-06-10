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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ThreadAffinity struct
type ThreadAffinity struct {
	*Thread_V2

	//
	Affinity uint32

	//
	Group uint16

	//
	Reserved uint16

	//
	ThreadId uint32
}

func NewThreadAffinityEx1(instance *cim.WmiInstance) (newInstance *ThreadAffinity, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ThreadAffinity{
		Thread_V2: tmp,
	}
	return
}

func NewThreadAffinityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ThreadAffinity, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ThreadAffinity{
		Thread_V2: tmp,
	}
	return
}

// SetAffinity sets the value of Affinity for the instance
func (instance *ThreadAffinity) SetPropertyAffinity(value uint32) (err error) {
	return instance.SetProperty("Affinity", (value))
}

// GetAffinity gets the value of Affinity for the instance
func (instance *ThreadAffinity) GetPropertyAffinity() (value uint32, err error) {
	retValue, err := instance.GetProperty("Affinity")
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

// SetGroup sets the value of Group for the instance
func (instance *ThreadAffinity) SetPropertyGroup(value uint16) (err error) {
	return instance.SetProperty("Group", (value))
}

// GetGroup gets the value of Group for the instance
func (instance *ThreadAffinity) GetPropertyGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *ThreadAffinity) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ThreadAffinity) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetThreadId sets the value of ThreadId for the instance
func (instance *ThreadAffinity) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *ThreadAffinity) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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
