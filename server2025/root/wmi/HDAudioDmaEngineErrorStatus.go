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

// HDAudioDmaEngineErrorStatus struct
type HDAudioDmaEngineErrorStatus struct {
	*cim.WmiInstance

	//
	DescriptorError uint32

	//
	FifoError uint32
}

func NewHDAudioDmaEngineErrorStatusEx1(instance *cim.WmiInstance) (newInstance *HDAudioDmaEngineErrorStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HDAudioDmaEngineErrorStatus{
		WmiInstance: tmp,
	}
	return
}

func NewHDAudioDmaEngineErrorStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HDAudioDmaEngineErrorStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HDAudioDmaEngineErrorStatus{
		WmiInstance: tmp,
	}
	return
}

// SetDescriptorError sets the value of DescriptorError for the instance
func (instance *HDAudioDmaEngineErrorStatus) SetPropertyDescriptorError(value uint32) (err error) {
	return instance.SetProperty("DescriptorError", (value))
}

// GetDescriptorError gets the value of DescriptorError for the instance
func (instance *HDAudioDmaEngineErrorStatus) GetPropertyDescriptorError() (value uint32, err error) {
	retValue, err := instance.GetProperty("DescriptorError")
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

// SetFifoError sets the value of FifoError for the instance
func (instance *HDAudioDmaEngineErrorStatus) SetPropertyFifoError(value uint32) (err error) {
	return instance.SetProperty("FifoError", (value))
}

// GetFifoError gets the value of FifoError for the instance
func (instance *HDAudioDmaEngineErrorStatus) GetPropertyFifoError() (value uint32, err error) {
	retValue, err := instance.GetProperty("FifoError")
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
