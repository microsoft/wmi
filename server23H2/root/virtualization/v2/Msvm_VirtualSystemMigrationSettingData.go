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

// Msvm_VirtualSystemMigrationSettingData struct
type Msvm_VirtualSystemMigrationSettingData struct {
	*CIM_VirtualSystemMigrationSettingData

	//
	AdvancedOptions string

	//
	AllowOverwriteExistingFile bool

	//
	AvoidRemovingVHDs bool

	//
	CancelIfBlackoutThresholdExceeded bool

	//
	CPUCappingMagnitude VirtualSystemMigrationSettingData_CPUCappingMagnitude

	//
	DestinationIPAddressList []string

	//
	DestinationPlannedVirtualSystemId string

	//
	EnableCompression bool

	//
	EnableVhdCompression bool

	//
	RemoveSourceUnmanagedVhds bool

	//
	RetainVhdCopiesOnSource bool

	//
	UnmanagedVhds []string
}

func NewMsvm_VirtualSystemMigrationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationSettingData{
		CIM_VirtualSystemMigrationSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationSettingData{
		CIM_VirtualSystemMigrationSettingData: tmp,
	}
	return
}

// SetAdvancedOptions sets the value of AdvancedOptions for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyAdvancedOptions(value string) (err error) {
	return instance.SetProperty("AdvancedOptions", (value))
}

// GetAdvancedOptions gets the value of AdvancedOptions for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyAdvancedOptions() (value string, err error) {
	retValue, err := instance.GetProperty("AdvancedOptions")
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

// SetAllowOverwriteExistingFile sets the value of AllowOverwriteExistingFile for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyAllowOverwriteExistingFile(value bool) (err error) {
	return instance.SetProperty("AllowOverwriteExistingFile", (value))
}

// GetAllowOverwriteExistingFile gets the value of AllowOverwriteExistingFile for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyAllowOverwriteExistingFile() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOverwriteExistingFile")
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

// SetAvoidRemovingVHDs sets the value of AvoidRemovingVHDs for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyAvoidRemovingVHDs(value bool) (err error) {
	return instance.SetProperty("AvoidRemovingVHDs", (value))
}

// GetAvoidRemovingVHDs gets the value of AvoidRemovingVHDs for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyAvoidRemovingVHDs() (value bool, err error) {
	retValue, err := instance.GetProperty("AvoidRemovingVHDs")
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

// SetCancelIfBlackoutThresholdExceeded sets the value of CancelIfBlackoutThresholdExceeded for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyCancelIfBlackoutThresholdExceeded(value bool) (err error) {
	return instance.SetProperty("CancelIfBlackoutThresholdExceeded", (value))
}

// GetCancelIfBlackoutThresholdExceeded gets the value of CancelIfBlackoutThresholdExceeded for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyCancelIfBlackoutThresholdExceeded() (value bool, err error) {
	retValue, err := instance.GetProperty("CancelIfBlackoutThresholdExceeded")
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

// SetCPUCappingMagnitude sets the value of CPUCappingMagnitude for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyCPUCappingMagnitude(value VirtualSystemMigrationSettingData_CPUCappingMagnitude) (err error) {
	return instance.SetProperty("CPUCappingMagnitude", (value))
}

// GetCPUCappingMagnitude gets the value of CPUCappingMagnitude for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyCPUCappingMagnitude() (value VirtualSystemMigrationSettingData_CPUCappingMagnitude, err error) {
	retValue, err := instance.GetProperty("CPUCappingMagnitude")
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

	value = VirtualSystemMigrationSettingData_CPUCappingMagnitude(valuetmp)

	return
}

// SetDestinationIPAddressList sets the value of DestinationIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyDestinationIPAddressList(value []string) (err error) {
	return instance.SetProperty("DestinationIPAddressList", (value))
}

// GetDestinationIPAddressList gets the value of DestinationIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyDestinationIPAddressList() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationIPAddressList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDestinationPlannedVirtualSystemId sets the value of DestinationPlannedVirtualSystemId for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyDestinationPlannedVirtualSystemId(value string) (err error) {
	return instance.SetProperty("DestinationPlannedVirtualSystemId", (value))
}

// GetDestinationPlannedVirtualSystemId gets the value of DestinationPlannedVirtualSystemId for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyDestinationPlannedVirtualSystemId() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPlannedVirtualSystemId")
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

// SetEnableCompression sets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyEnableCompression(value bool) (err error) {
	return instance.SetProperty("EnableCompression", (value))
}

// GetEnableCompression gets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyEnableCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableCompression")
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

// SetEnableVhdCompression sets the value of EnableVhdCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyEnableVhdCompression(value bool) (err error) {
	return instance.SetProperty("EnableVhdCompression", (value))
}

// GetEnableVhdCompression gets the value of EnableVhdCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyEnableVhdCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableVhdCompression")
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

// SetRemoveSourceUnmanagedVhds sets the value of RemoveSourceUnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyRemoveSourceUnmanagedVhds(value bool) (err error) {
	return instance.SetProperty("RemoveSourceUnmanagedVhds", (value))
}

// GetRemoveSourceUnmanagedVhds gets the value of RemoveSourceUnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyRemoveSourceUnmanagedVhds() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoveSourceUnmanagedVhds")
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

// SetRetainVhdCopiesOnSource sets the value of RetainVhdCopiesOnSource for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyRetainVhdCopiesOnSource(value bool) (err error) {
	return instance.SetProperty("RetainVhdCopiesOnSource", (value))
}

// GetRetainVhdCopiesOnSource gets the value of RetainVhdCopiesOnSource for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyRetainVhdCopiesOnSource() (value bool, err error) {
	retValue, err := instance.GetProperty("RetainVhdCopiesOnSource")
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

// SetUnmanagedVhds sets the value of UnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyUnmanagedVhds(value []string) (err error) {
	return instance.SetProperty("UnmanagedVhds", (value))
}

// GetUnmanagedVhds gets the value of UnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyUnmanagedVhds() (value []string, err error) {
	retValue, err := instance.GetProperty("UnmanagedVhds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
func (instance *Msvm_VirtualSystemMigrationSettingData) GetRelatedVirtualSystemMigrationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationCapabilities")
}
