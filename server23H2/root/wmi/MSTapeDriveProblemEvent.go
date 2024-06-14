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

// MSTapeDriveProblemEvent struct
type MSTapeDriveProblemEvent struct {
	*WMIEvent

	//
	Active bool

	//
	DriveProblemType uint32

	//
	InstanceName string

	//
	TapeData []uint8
}

func NewMSTapeDriveProblemEventEx1(instance *cim.WmiInstance) (newInstance *MSTapeDriveProblemEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSTapeDriveProblemEvent{
		WMIEvent: tmp,
	}
	return
}

func NewMSTapeDriveProblemEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSTapeDriveProblemEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSTapeDriveProblemEvent{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSTapeDriveProblemEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeDriveProblemEvent) GetPropertyActive() (value bool, err error) {
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

// SetDriveProblemType sets the value of DriveProblemType for the instance
func (instance *MSTapeDriveProblemEvent) SetPropertyDriveProblemType(value uint32) (err error) {
	return instance.SetProperty("DriveProblemType", (value))
}

// GetDriveProblemType gets the value of DriveProblemType for the instance
func (instance *MSTapeDriveProblemEvent) GetPropertyDriveProblemType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveProblemType")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSTapeDriveProblemEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeDriveProblemEvent) GetPropertyInstanceName() (value string, err error) {
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

// SetTapeData sets the value of TapeData for the instance
func (instance *MSTapeDriveProblemEvent) SetPropertyTapeData(value []uint8) (err error) {
	return instance.SetProperty("TapeData", (value))
}

// GetTapeData gets the value of TapeData for the instance
func (instance *MSTapeDriveProblemEvent) GetPropertyTapeData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("TapeData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
