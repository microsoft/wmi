// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// HDAudioControllerErrorStatus struct
type HDAudioControllerErrorStatus struct {
	*cim.WmiInstance

	//
	CorbMemoryError uint32

	//
	RirbFifoOverrun uint32
}

func NewHDAudioControllerErrorStatusEx1(instance *cim.WmiInstance) (newInstance *HDAudioControllerErrorStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HDAudioControllerErrorStatus{
		WmiInstance: tmp,
	}
	return
}

func NewHDAudioControllerErrorStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HDAudioControllerErrorStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HDAudioControllerErrorStatus{
		WmiInstance: tmp,
	}
	return
}

// SetCorbMemoryError sets the value of CorbMemoryError for the instance
func (instance *HDAudioControllerErrorStatus) SetPropertyCorbMemoryError(value uint32) (err error) {
	return instance.SetProperty("CorbMemoryError", (value))
}

// GetCorbMemoryError gets the value of CorbMemoryError for the instance
func (instance *HDAudioControllerErrorStatus) GetPropertyCorbMemoryError() (value uint32, err error) {
	retValue, err := instance.GetProperty("CorbMemoryError")
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

// SetRirbFifoOverrun sets the value of RirbFifoOverrun for the instance
func (instance *HDAudioControllerErrorStatus) SetPropertyRirbFifoOverrun(value uint32) (err error) {
	return instance.SetProperty("RirbFifoOverrun", (value))
}

// GetRirbFifoOverrun gets the value of RirbFifoOverrun for the instance
func (instance *HDAudioControllerErrorStatus) GetPropertyRirbFifoOverrun() (value uint32, err error) {
	retValue, err := instance.GetProperty("RirbFifoOverrun")
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
