// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_SoftwareElement struct
type CIM_SoftwareElement struct {
	*CIM_LogicalElement

	// The internal identifier for this compilation of SoftwareElement.
	BuildNumber string

	// The code set used by this SoftwareElement. It defines the bit patterns that a system uses to identify characters. ISO defines various code sets such as UTF-8 and ISO8859-1.
	CodeSet string

	// The manufacturer's identifier for this SoftwareElement. Often this will be a stock keeping unit (SKU) or a part number.
	IdentificationCode string

	// The value of this property identifies the language edition of this SoftwareElement. The language codes defined in ISO 639 should be used. Where the element represents a multi-lingual or international version, the string "Multilingual" should be used.
	LanguageEdition string

	// Manufacturer of this SoftwareElement.
	Manufacturer string

	// The OtherTargetOS property records the manufacturer and operating system type for a SoftwareElement when the TargetOperatingSystem property has a value of 1 ("Other"). For all other values of TargetOperatingSystem, the OtherTargetOS property is NULL.
	OtherTargetOS string

	// The assigned serial number of this SoftwareElement.
	SerialNumber string

	// This is an identifier for the SoftwareElement and is designed to be used in conjunction with other keys to create a unique representation of the element.
	SoftwareElementID string

	// The SoftwareElementState is defined in this model to identify various states of a SoftwareElement's life cycle.
	///- A SoftwareElement in the deployable state describes the details necessary to successfully distribute it and the details (Checks and Actions) required to move it to the installable state (i.e, the next state).
	///- A SoftwareElement in the installable state describes the details necessary to successfully install it and the details (Checks and Actions) required to create an element in the executable state (i.e., the next state).
	///- A SoftwareElement in the executable state describes the details necessary to successfully start it and the details (Checks and Actions) required to move it to the running state (i.e., the next state).
	///- A SoftwareElement in the running state describes the details necessary to manage the started element.
	SoftwareElementState SoftwareElement_SoftwareElementState

	// The TargetOperatingSystem property specifies the element's operating system environment. The value of this property does not ensure that it is binary executable. Two other pieces of information are needed. First, the version of the OS needs to be specified using the class, CIM_OSVersion Check. The second piece of information is the architecture that the OS runs on. This information is verified using CIM_ArchitectureCheck. The combination of these constructs clearly identifies the level of OS required for a particular SoftwareElement.
	TargetOperatingSystem SoftwareElement_TargetOperatingSystem

	// Software Version should be in the form <Major>.<Minor>.<Revision> or <Major>.<Minor><letter><revision>.
	Version string
}

func NewCIM_SoftwareElementEx1(instance *cim.WmiInstance) (newInstance *CIM_SoftwareElement, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareElement{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_SoftwareElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SoftwareElement, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareElement{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *CIM_SoftwareElement) SetPropertyBuildNumber(value string) (err error) {
	return instance.SetProperty("BuildNumber", (value))
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *CIM_SoftwareElement) GetPropertyBuildNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
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

// SetCodeSet sets the value of CodeSet for the instance
func (instance *CIM_SoftwareElement) SetPropertyCodeSet(value string) (err error) {
	return instance.SetProperty("CodeSet", (value))
}

// GetCodeSet gets the value of CodeSet for the instance
func (instance *CIM_SoftwareElement) GetPropertyCodeSet() (value string, err error) {
	retValue, err := instance.GetProperty("CodeSet")
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

// SetIdentificationCode sets the value of IdentificationCode for the instance
func (instance *CIM_SoftwareElement) SetPropertyIdentificationCode(value string) (err error) {
	return instance.SetProperty("IdentificationCode", (value))
}

// GetIdentificationCode gets the value of IdentificationCode for the instance
func (instance *CIM_SoftwareElement) GetPropertyIdentificationCode() (value string, err error) {
	retValue, err := instance.GetProperty("IdentificationCode")
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

// SetLanguageEdition sets the value of LanguageEdition for the instance
func (instance *CIM_SoftwareElement) SetPropertyLanguageEdition(value string) (err error) {
	return instance.SetProperty("LanguageEdition", (value))
}

// GetLanguageEdition gets the value of LanguageEdition for the instance
func (instance *CIM_SoftwareElement) GetPropertyLanguageEdition() (value string, err error) {
	retValue, err := instance.GetProperty("LanguageEdition")
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
func (instance *CIM_SoftwareElement) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_SoftwareElement) GetPropertyManufacturer() (value string, err error) {
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

// SetOtherTargetOS sets the value of OtherTargetOS for the instance
func (instance *CIM_SoftwareElement) SetPropertyOtherTargetOS(value string) (err error) {
	return instance.SetProperty("OtherTargetOS", (value))
}

// GetOtherTargetOS gets the value of OtherTargetOS for the instance
func (instance *CIM_SoftwareElement) GetPropertyOtherTargetOS() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTargetOS")
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
func (instance *CIM_SoftwareElement) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_SoftwareElement) GetPropertySerialNumber() (value string, err error) {
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

// SetSoftwareElementID sets the value of SoftwareElementID for the instance
func (instance *CIM_SoftwareElement) SetPropertySoftwareElementID(value string) (err error) {
	return instance.SetProperty("SoftwareElementID", (value))
}

// GetSoftwareElementID gets the value of SoftwareElementID for the instance
func (instance *CIM_SoftwareElement) GetPropertySoftwareElementID() (value string, err error) {
	retValue, err := instance.GetProperty("SoftwareElementID")
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

// SetSoftwareElementState sets the value of SoftwareElementState for the instance
func (instance *CIM_SoftwareElement) SetPropertySoftwareElementState(value SoftwareElement_SoftwareElementState) (err error) {
	return instance.SetProperty("SoftwareElementState", (value))
}

// GetSoftwareElementState gets the value of SoftwareElementState for the instance
func (instance *CIM_SoftwareElement) GetPropertySoftwareElementState() (value SoftwareElement_SoftwareElementState, err error) {
	retValue, err := instance.GetProperty("SoftwareElementState")
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

	value = SoftwareElement_SoftwareElementState(valuetmp)

	return
}

// SetTargetOperatingSystem sets the value of TargetOperatingSystem for the instance
func (instance *CIM_SoftwareElement) SetPropertyTargetOperatingSystem(value SoftwareElement_TargetOperatingSystem) (err error) {
	return instance.SetProperty("TargetOperatingSystem", (value))
}

// GetTargetOperatingSystem gets the value of TargetOperatingSystem for the instance
func (instance *CIM_SoftwareElement) GetPropertyTargetOperatingSystem() (value SoftwareElement_TargetOperatingSystem, err error) {
	retValue, err := instance.GetProperty("TargetOperatingSystem")
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

	value = SoftwareElement_TargetOperatingSystem(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *CIM_SoftwareElement) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *CIM_SoftwareElement) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
