// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ControlledBy struct
type CIM_ControlledBy struct {
	*CIM_DeviceConnection

	// This property describes the accessibility of the device through the antecedent controller.
	AccessMode ControlledBy_AccessMode

	// The property describes the priority given to accesses of the device through this controller. The highest priority path will have the lowest value for this parameter.
	AccessPriority uint16

	// The State property indicates whether the Controller is actively commanding or accessing the Device (value=1) or not (value=2). Also, the value, "Unknown" (0), can be defined. This information is necessary when a LogicalDevice can be commanded by, or accessed through, multiple Controllers.
	AccessState ControlledBy_AccessState

	// Address of associated Device in context of the antecedent Controller.
	DeviceNumber string

	// Number of hard resets issued by the Controller. A hard reset returns the Device to its initialization or boot-up state. All internal Device state information and data are lost.
	NumberOfHardResets uint32

	// Number of soft resets issued by the Controller. A soft reset does not completely clear current Device state or data. Exact semantics are dependent on the Device and on the protocols and mechanisms used to communicate with the Device.
	NumberOfSoftResets uint32

	// The time that the downstream Device was last reset by the Controller.
	TimeOfDeviceReset string
}

func NewCIM_ControlledByEx1(instance *cim.WmiInstance) (newInstance *CIM_ControlledBy, err error) {
	tmp, err := NewCIM_DeviceConnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ControlledBy{
		CIM_DeviceConnection: tmp,
	}
	return
}

func NewCIM_ControlledByEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ControlledBy, err error) {
	tmp, err := NewCIM_DeviceConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ControlledBy{
		CIM_DeviceConnection: tmp,
	}
	return
}

// SetAccessMode sets the value of AccessMode for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessMode(value ControlledBy_AccessMode) (err error) {
	return instance.SetProperty("AccessMode", (value))
}

// GetAccessMode gets the value of AccessMode for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessMode() (value ControlledBy_AccessMode, err error) {
	retValue, err := instance.GetProperty("AccessMode")
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

	value = ControlledBy_AccessMode(valuetmp)

	return
}

// SetAccessPriority sets the value of AccessPriority for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessPriority(value uint16) (err error) {
	return instance.SetProperty("AccessPriority", (value))
}

// GetAccessPriority gets the value of AccessPriority for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessPriority")
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

// SetAccessState sets the value of AccessState for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessState(value ControlledBy_AccessState) (err error) {
	return instance.SetProperty("AccessState", (value))
}

// GetAccessState gets the value of AccessState for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessState() (value ControlledBy_AccessState, err error) {
	retValue, err := instance.GetProperty("AccessState")
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

	value = ControlledBy_AccessState(valuetmp)

	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *CIM_ControlledBy) SetPropertyDeviceNumber(value string) (err error) {
	return instance.SetProperty("DeviceNumber", (value))
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *CIM_ControlledBy) GetPropertyDeviceNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
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

// SetNumberOfHardResets sets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfHardResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfHardResets", (value))
}

// GetNumberOfHardResets gets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfHardResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfHardResets")
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

// SetNumberOfSoftResets sets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfSoftResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfSoftResets", (value))
}

// GetNumberOfSoftResets gets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfSoftResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSoftResets")
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

// SetTimeOfDeviceReset sets the value of TimeOfDeviceReset for the instance
func (instance *CIM_ControlledBy) SetPropertyTimeOfDeviceReset(value string) (err error) {
	return instance.SetProperty("TimeOfDeviceReset", (value))
}

// GetTimeOfDeviceReset gets the value of TimeOfDeviceReset for the instance
func (instance *CIM_ControlledBy) GetPropertyTimeOfDeviceReset() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfDeviceReset")
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
