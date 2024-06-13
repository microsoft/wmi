// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_MountedStorageImage struct
type Msvm_MountedStorageImage struct {
	*CIM_LogicalElement

	//
	Access uint16

	//
	Lun uint8

	//
	PathId uint8

	//
	PnpDevicePath string

	//
	PortNumber uint8

	//
	TargetId uint8

	//
	Type uint16
}

func NewMsvm_MountedStorageImageEx1(instance *cim.WmiInstance) (newInstance *Msvm_MountedStorageImage, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MountedStorageImage{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMsvm_MountedStorageImageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MountedStorageImage, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MountedStorageImage{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetAccess sets the value of Access for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyAccess(value uint16) (err error) {
	return instance.SetProperty("Access", (value))
}

// GetAccess gets the value of Access for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyAccess() (value uint16, err error) {
	retValue, err := instance.GetProperty("Access")
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

// SetLun sets the value of Lun for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyLun(value uint8) (err error) {
	return instance.SetProperty("Lun", (value))
}

// GetLun gets the value of Lun for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyLun() (value uint8, err error) {
	retValue, err := instance.GetProperty("Lun")
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

// SetPathId sets the value of PathId for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPathId(value uint8) (err error) {
	return instance.SetProperty("PathId", (value))
}

// GetPathId gets the value of PathId for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPathId() (value uint8, err error) {
	retValue, err := instance.GetProperty("PathId")
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

// SetPnpDevicePath sets the value of PnpDevicePath for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPnpDevicePath(value string) (err error) {
	return instance.SetProperty("PnpDevicePath", (value))
}

// GetPnpDevicePath gets the value of PnpDevicePath for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPnpDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("PnpDevicePath")
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

// SetPortNumber sets the value of PortNumber for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPortNumber(value uint8) (err error) {
	return instance.SetProperty("PortNumber", (value))
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPortNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PortNumber")
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

// SetTargetId sets the value of TargetId for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyTargetId(value uint8) (err error) {
	return instance.SetProperty("TargetId", (value))
}

// GetTargetId gets the value of TargetId for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyTargetId() (value uint8, err error) {
	retValue, err := instance.GetProperty("TargetId")
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

// SetType sets the value of Type for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_MountedStorageImage) DetachVirtualHardDisk() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DetachVirtualHardDisk")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
