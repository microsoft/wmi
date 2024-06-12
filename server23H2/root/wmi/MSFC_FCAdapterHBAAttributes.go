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

// MSFC_FCAdapterHBAAttributes struct
type MSFC_FCAdapterHBAAttributes struct {
	*cim.WmiInstance

	//
	Active bool

	//
	DriverName string

	//
	DriverVersion string

	//
	FirmwareVersion string

	//
	HardwareVersion string

	//
	HBAStatus uint32

	//
	InstanceName string

	//
	Manufacturer string

	//
	MfgDomain string

	//
	Model string

	//
	ModelDescription string

	//
	NodeSymbolicName string

	//
	NodeWWN []uint8

	//
	NumberOfPorts uint32

	//
	OptionROMVersion string

	//
	SerialNumber string

	//
	UniqueAdapterId uint64

	//
	VendorSpecificID uint32
}

func NewMSFC_FCAdapterHBAAttributesEx1(instance *cim.WmiInstance) (newInstance *MSFC_FCAdapterHBAAttributes, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FCAdapterHBAAttributes{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FCAdapterHBAAttributesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FCAdapterHBAAttributes, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FCAdapterHBAAttributes{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyActive() (value bool, err error) {
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

// SetDriverName sets the value of DriverName for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyDriverName(value string) (err error) {
	return instance.SetProperty("DriverName", (value))
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyDriverName() (value string, err error) {
	retValue, err := instance.GetProperty("DriverName")
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

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyDriverVersion(value string) (err error) {
	return instance.SetProperty("DriverVersion", (value))
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyDriverVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
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

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
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

// SetHardwareVersion sets the value of HardwareVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyHardwareVersion(value string) (err error) {
	return instance.SetProperty("HardwareVersion", (value))
}

// GetHardwareVersion gets the value of HardwareVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyHardwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareVersion")
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

// SetHBAStatus sets the value of HBAStatus for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyHBAStatus(value uint32) (err error) {
	return instance.SetProperty("HBAStatus", (value))
}

// GetHBAStatus gets the value of HBAStatus for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyHBAStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HBAStatus")
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
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyInstanceName() (value string, err error) {
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetMfgDomain sets the value of MfgDomain for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyMfgDomain(value string) (err error) {
	return instance.SetProperty("MfgDomain", (value))
}

// GetMfgDomain gets the value of MfgDomain for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyMfgDomain() (value string, err error) {
	retValue, err := instance.GetProperty("MfgDomain")
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

// SetModel sets the value of Model for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
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

// SetModelDescription sets the value of ModelDescription for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyModelDescription(value string) (err error) {
	return instance.SetProperty("ModelDescription", (value))
}

// GetModelDescription gets the value of ModelDescription for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyModelDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ModelDescription")
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

// SetNodeSymbolicName sets the value of NodeSymbolicName for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyNodeSymbolicName(value string) (err error) {
	return instance.SetProperty("NodeSymbolicName", (value))
}

// GetNodeSymbolicName gets the value of NodeSymbolicName for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyNodeSymbolicName() (value string, err error) {
	retValue, err := instance.GetProperty("NodeSymbolicName")
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

// SetNodeWWN sets the value of NodeWWN for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyNodeWWN(value []uint8) (err error) {
	return instance.SetProperty("NodeWWN", (value))
}

// GetNodeWWN gets the value of NodeWWN for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyNodeWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("NodeWWN")
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

// SetNumberOfPorts sets the value of NumberOfPorts for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyNumberOfPorts(value uint32) (err error) {
	return instance.SetProperty("NumberOfPorts", (value))
}

// GetNumberOfPorts gets the value of NumberOfPorts for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyNumberOfPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPorts")
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

// SetOptionROMVersion sets the value of OptionROMVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyOptionROMVersion(value string) (err error) {
	return instance.SetProperty("OptionROMVersion", (value))
}

// GetOptionROMVersion gets the value of OptionROMVersion for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyOptionROMVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OptionROMVersion")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyUniqueAdapterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueAdapterId")
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

// SetVendorSpecificID sets the value of VendorSpecificID for the instance
func (instance *MSFC_FCAdapterHBAAttributes) SetPropertyVendorSpecificID(value uint32) (err error) {
	return instance.SetProperty("VendorSpecificID", (value))
}

// GetVendorSpecificID gets the value of VendorSpecificID for the instance
func (instance *MSFC_FCAdapterHBAAttributes) GetPropertyVendorSpecificID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VendorSpecificID")
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
