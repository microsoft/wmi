// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSParallel_AllocFreeCounts struct
type MSParallel_AllocFreeCounts struct {
	*MSParallel

	//
	Active bool

	//
	InstanceName string

	//
	PortAllocates uint32

	//
	PortFrees uint32
}

func NewMSParallel_AllocFreeCountsEx1(instance *cim.WmiInstance) (newInstance *MSParallel_AllocFreeCounts, err error) {
	tmp, err := NewMSParallelEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSParallel_AllocFreeCounts{
		MSParallel: tmp,
	}
	return
}

func NewMSParallel_AllocFreeCountsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSParallel_AllocFreeCounts, err error) {
	tmp, err := NewMSParallelEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSParallel_AllocFreeCounts{
		MSParallel: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSParallel_AllocFreeCounts) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSParallel_AllocFreeCounts) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSParallel_AllocFreeCounts) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSParallel_AllocFreeCounts) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetPortAllocates sets the value of PortAllocates for the instance
func (instance *MSParallel_AllocFreeCounts) SetPropertyPortAllocates(value uint32) (err error) {
	return instance.SetProperty("PortAllocates", (value))
}

// GetPortAllocates gets the value of PortAllocates for the instance
func (instance *MSParallel_AllocFreeCounts) GetPropertyPortAllocates() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortAllocates")
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

// SetPortFrees sets the value of PortFrees for the instance
func (instance *MSParallel_AllocFreeCounts) SetPropertyPortFrees(value uint32) (err error) {
	return instance.SetProperty("PortFrees", (value))
}

// GetPortFrees gets the value of PortFrees for the instance
func (instance *MSParallel_AllocFreeCounts) GetPropertyPortFrees() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortFrees")
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
