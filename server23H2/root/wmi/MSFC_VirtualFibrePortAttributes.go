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

// MSFC_VirtualFibrePortAttributes struct
type MSFC_VirtualFibrePortAttributes struct {
	*cim.WmiInstance

	//
	FabricWWN []uint8

	//
	FCId uint32

	//
	Status VirtualFibrePortAttributes_Status

	//
	Tag []uint8

	//
	VirtualName []uint16

	//
	WWNN []uint8

	//
	WWPN []uint8
}

func NewMSFC_VirtualFibrePortAttributesEx1(instance *cim.WmiInstance) (newInstance *MSFC_VirtualFibrePortAttributes, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_VirtualFibrePortAttributes{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_VirtualFibrePortAttributesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_VirtualFibrePortAttributes, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_VirtualFibrePortAttributes{
		WmiInstance: tmp,
	}
	return
}

// SetFabricWWN sets the value of FabricWWN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyFabricWWN(value []uint8) (err error) {
	return instance.SetProperty("FabricWWN", (value))
}

// GetFabricWWN gets the value of FabricWWN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyFabricWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FabricWWN")
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

// SetFCId sets the value of FCId for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyFCId(value uint32) (err error) {
	return instance.SetProperty("FCId", (value))
}

// GetFCId gets the value of FCId for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyFCId() (value uint32, err error) {
	retValue, err := instance.GetProperty("FCId")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyStatus(value VirtualFibrePortAttributes_Status) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyStatus() (value VirtualFibrePortAttributes_Status, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualFibrePortAttributes_Status(valuetmp)

	return
}

// SetTag sets the value of Tag for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyTag(value []uint8) (err error) {
	return instance.SetProperty("Tag", (value))
}

// GetTag gets the value of Tag for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyTag() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Tag")
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

// SetVirtualName sets the value of VirtualName for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyVirtualName(value []uint16) (err error) {
	return instance.SetProperty("VirtualName", (value))
}

// GetVirtualName gets the value of VirtualName for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyVirtualName() (value []uint16, err error) {
	retValue, err := instance.GetProperty("VirtualName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetWWNN sets the value of WWNN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyWWNN(value []uint8) (err error) {
	return instance.SetProperty("WWNN", (value))
}

// GetWWNN gets the value of WWNN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyWWNN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWNN")
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

// SetWWPN sets the value of WWPN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) SetPropertyWWPN(value []uint8) (err error) {
	return instance.SetProperty("WWPN", (value))
}

// GetWWPN gets the value of WWPN for the instance
func (instance *MSFC_VirtualFibrePortAttributes) GetPropertyWWPN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWPN")
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
