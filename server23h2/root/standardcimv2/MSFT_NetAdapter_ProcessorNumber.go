// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_ProcessorNumber struct
type MSFT_NetAdapter_ProcessorNumber struct {
	*cim.WmiInstance

	//
	ProcessorGroup uint16

	//
	ProcessorNumber uint8
}

func NewMSFT_NetAdapter_ProcessorNumberEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_ProcessorNumber, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_ProcessorNumber{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_ProcessorNumberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_ProcessorNumber, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_ProcessorNumber{
		WmiInstance: tmp,
	}
	return
}

// SetProcessorGroup sets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_ProcessorNumber) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", (value))
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_ProcessorNumber) GetPropertyProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorGroup")
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

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_ProcessorNumber) SetPropertyProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("ProcessorNumber", (value))
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_ProcessorNumber) GetPropertyProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("ProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
