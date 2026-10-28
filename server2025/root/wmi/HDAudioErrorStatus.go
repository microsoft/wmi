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

// HDAudioErrorStatus struct
type HDAudioErrorStatus struct {
	*cim.WmiInstance

	//
	Active bool

	//
	ControllerErrorStatus HDAudioControllerErrorStatus

	//
	EngineErrorStatus []HDAudioDmaEngineErrorStatus

	//
	InstanceName string

	//
	NumEngines uint32
}

func NewHDAudioErrorStatusEx1(instance *cim.WmiInstance) (newInstance *HDAudioErrorStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HDAudioErrorStatus{
		WmiInstance: tmp,
	}
	return
}

func NewHDAudioErrorStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HDAudioErrorStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HDAudioErrorStatus{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *HDAudioErrorStatus) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *HDAudioErrorStatus) GetPropertyActive() (value bool, err error) {
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

// SetControllerErrorStatus sets the value of ControllerErrorStatus for the instance
func (instance *HDAudioErrorStatus) SetPropertyControllerErrorStatus(value HDAudioControllerErrorStatus) (err error) {
	return instance.SetProperty("ControllerErrorStatus", (value))
}

// GetControllerErrorStatus gets the value of ControllerErrorStatus for the instance
func (instance *HDAudioErrorStatus) GetPropertyControllerErrorStatus() (value HDAudioControllerErrorStatus, err error) {
	retValue, err := instance.GetProperty("ControllerErrorStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HDAudioControllerErrorStatus)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HDAudioControllerErrorStatus is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HDAudioControllerErrorStatus(valuetmp)

	return
}

// SetEngineErrorStatus sets the value of EngineErrorStatus for the instance
func (instance *HDAudioErrorStatus) SetPropertyEngineErrorStatus(value []HDAudioDmaEngineErrorStatus) (err error) {
	return instance.SetProperty("EngineErrorStatus", (value))
}

// GetEngineErrorStatus gets the value of EngineErrorStatus for the instance
func (instance *HDAudioErrorStatus) GetPropertyEngineErrorStatus() (value []HDAudioDmaEngineErrorStatus, err error) {
	retValue, err := instance.GetProperty("EngineErrorStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(HDAudioDmaEngineErrorStatus)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " HDAudioDmaEngineErrorStatus is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, HDAudioDmaEngineErrorStatus(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *HDAudioErrorStatus) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *HDAudioErrorStatus) GetPropertyInstanceName() (value string, err error) {
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

// SetNumEngines sets the value of NumEngines for the instance
func (instance *HDAudioErrorStatus) SetPropertyNumEngines(value uint32) (err error) {
	return instance.SetProperty("NumEngines", (value))
}

// GetNumEngines gets the value of NumEngines for the instance
func (instance *HDAudioErrorStatus) GetPropertyNumEngines() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumEngines")
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
