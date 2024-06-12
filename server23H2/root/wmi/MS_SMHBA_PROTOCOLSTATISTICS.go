// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MS_SMHBA_PROTOCOLSTATISTICS struct
type MS_SMHBA_PROTOCOLSTATISTICS struct {
	*cim.WmiInstance

	//
	ControlRequests int64

	//
	InputMegabytes int64

	//
	InputRequests int64

	//
	OutputMegabytes int64

	//
	OutputRequests int64

	//
	SecondsSinceLastReset int64
}

func NewMS_SMHBA_PROTOCOLSTATISTICSEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_PROTOCOLSTATISTICS, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PROTOCOLSTATISTICS{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_PROTOCOLSTATISTICSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_PROTOCOLSTATISTICS, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PROTOCOLSTATISTICS{
		WmiInstance: tmp,
	}
	return
}

// SetControlRequests sets the value of ControlRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertyControlRequests(value int64) (err error) {
	return instance.SetProperty("ControlRequests", (value))
}

// GetControlRequests gets the value of ControlRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertyControlRequests() (value int64, err error) {
	retValue, err := instance.GetProperty("ControlRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetInputMegabytes sets the value of InputMegabytes for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertyInputMegabytes(value int64) (err error) {
	return instance.SetProperty("InputMegabytes", (value))
}

// GetInputMegabytes gets the value of InputMegabytes for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertyInputMegabytes() (value int64, err error) {
	retValue, err := instance.GetProperty("InputMegabytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetInputRequests sets the value of InputRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertyInputRequests(value int64) (err error) {
	return instance.SetProperty("InputRequests", (value))
}

// GetInputRequests gets the value of InputRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertyInputRequests() (value int64, err error) {
	retValue, err := instance.GetProperty("InputRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetOutputMegabytes sets the value of OutputMegabytes for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertyOutputMegabytes(value int64) (err error) {
	return instance.SetProperty("OutputMegabytes", (value))
}

// GetOutputMegabytes gets the value of OutputMegabytes for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertyOutputMegabytes() (value int64, err error) {
	retValue, err := instance.GetProperty("OutputMegabytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetOutputRequests sets the value of OutputRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertyOutputRequests(value int64) (err error) {
	return instance.SetProperty("OutputRequests", (value))
}

// GetOutputRequests gets the value of OutputRequests for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertyOutputRequests() (value int64, err error) {
	retValue, err := instance.GetProperty("OutputRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetSecondsSinceLastReset sets the value of SecondsSinceLastReset for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) SetPropertySecondsSinceLastReset(value int64) (err error) {
	return instance.SetProperty("SecondsSinceLastReset", (value))
}

// GetSecondsSinceLastReset gets the value of SecondsSinceLastReset for the instance
func (instance *MS_SMHBA_PROTOCOLSTATISTICS) GetPropertySecondsSinceLastReset() (value int64, err error) {
	retValue, err := instance.GetProperty("SecondsSinceLastReset")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}
