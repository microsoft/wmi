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

// MSFC_FC4STATISTICS struct
type MSFC_FC4STATISTICS struct {
	*cim.WmiInstance

	//
	ControlRequests uint64

	//
	InputMegabytes uint64

	//
	InputRequests uint64

	//
	OutputMegabytes uint64

	//
	OutputRequests uint64
}

func NewMSFC_FC4STATISTICSEx1(instance *cim.WmiInstance) (newInstance *MSFC_FC4STATISTICS, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FC4STATISTICS{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FC4STATISTICSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FC4STATISTICS, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FC4STATISTICS{
		WmiInstance: tmp,
	}
	return
}

// SetControlRequests sets the value of ControlRequests for the instance
func (instance *MSFC_FC4STATISTICS) SetPropertyControlRequests(value uint64) (err error) {
	return instance.SetProperty("ControlRequests", (value))
}

// GetControlRequests gets the value of ControlRequests for the instance
func (instance *MSFC_FC4STATISTICS) GetPropertyControlRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlRequests")
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

// SetInputMegabytes sets the value of InputMegabytes for the instance
func (instance *MSFC_FC4STATISTICS) SetPropertyInputMegabytes(value uint64) (err error) {
	return instance.SetProperty("InputMegabytes", (value))
}

// GetInputMegabytes gets the value of InputMegabytes for the instance
func (instance *MSFC_FC4STATISTICS) GetPropertyInputMegabytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("InputMegabytes")
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

// SetInputRequests sets the value of InputRequests for the instance
func (instance *MSFC_FC4STATISTICS) SetPropertyInputRequests(value uint64) (err error) {
	return instance.SetProperty("InputRequests", (value))
}

// GetInputRequests gets the value of InputRequests for the instance
func (instance *MSFC_FC4STATISTICS) GetPropertyInputRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("InputRequests")
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

// SetOutputMegabytes sets the value of OutputMegabytes for the instance
func (instance *MSFC_FC4STATISTICS) SetPropertyOutputMegabytes(value uint64) (err error) {
	return instance.SetProperty("OutputMegabytes", (value))
}

// GetOutputMegabytes gets the value of OutputMegabytes for the instance
func (instance *MSFC_FC4STATISTICS) GetPropertyOutputMegabytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutputMegabytes")
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

// SetOutputRequests sets the value of OutputRequests for the instance
func (instance *MSFC_FC4STATISTICS) SetPropertyOutputRequests(value uint64) (err error) {
	return instance.SetProperty("OutputRequests", (value))
}

// GetOutputRequests gets the value of OutputRequests for the instance
func (instance *MSFC_FC4STATISTICS) GetPropertyOutputRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutputRequests")
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
