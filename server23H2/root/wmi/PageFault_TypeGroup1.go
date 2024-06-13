// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PageFault_TypeGroup1 struct
type PageFault_TypeGroup1 struct {
	*PageFault_V2

	//
	ProgramCounter uint32

	//
	VirtualAddress uint32
}

func NewPageFault_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *PageFault_TypeGroup1, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_TypeGroup1{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_TypeGroup1, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_TypeGroup1{
		PageFault_V2: tmp,
	}
	return
}

// SetProgramCounter sets the value of ProgramCounter for the instance
func (instance *PageFault_TypeGroup1) SetPropertyProgramCounter(value uint32) (err error) {
	return instance.SetProperty("ProgramCounter", (value))
}

// GetProgramCounter gets the value of ProgramCounter for the instance
func (instance *PageFault_TypeGroup1) GetPropertyProgramCounter() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProgramCounter")
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

// SetVirtualAddress sets the value of VirtualAddress for the instance
func (instance *PageFault_TypeGroup1) SetPropertyVirtualAddress(value uint32) (err error) {
	return instance.SetProperty("VirtualAddress", (value))
}

// GetVirtualAddress gets the value of VirtualAddress for the instance
func (instance *PageFault_TypeGroup1) GetPropertyVirtualAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualAddress")
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
