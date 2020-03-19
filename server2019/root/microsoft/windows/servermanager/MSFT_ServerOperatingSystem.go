// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerOperatingSystem struct
type MSFT_ServerOperatingSystem struct {
	*cim.WmiInstance

	//
	Architecture uint8

	//
	BuildNumber uint32

	//
	Language string

	//
	MajorVersion uint32

	//
	MinorVersion uint32

	//
	Name string

	//
	SKU string

	//
	SKUId uint32

	//
	SPMajorVersion uint16

	//
	SPMinorVersion uint16
}

func NewMSFT_ServerOperatingSystemEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerOperatingSystem, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerOperatingSystem{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerOperatingSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerOperatingSystem, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerOperatingSystem{
		WmiInstance: tmp,
	}
	return
}

// SetArchitecture sets the value of Architecture for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyArchitecture(value uint8) (err error) {
	return instance.SetProperty("Architecture", value)
}

// GetArchitecture gets the value of Architecture for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyArchitecture() (value uint8, err error) {
	retValue, err := instance.GetProperty("Architecture")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyBuildNumber(value uint32) (err error) {
	return instance.SetProperty("BuildNumber", value)
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyBuildNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanguage sets the value of Language for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyLanguage(value string) (err error) {
	return instance.SetProperty("Language", value)
}

// GetLanguage gets the value of Language for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("Language")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMajorVersion sets the value of MajorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyMajorVersion(value uint32) (err error) {
	return instance.SetProperty("MajorVersion", value)
}

// GetMajorVersion gets the value of MajorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyMajorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MajorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinorVersion sets the value of MinorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyMinorVersion(value uint32) (err error) {
	return instance.SetProperty("MinorVersion", value)
}

// GetMinorVersion gets the value of MinorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyMinorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSKU sets the value of SKU for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertySKU(value string) (err error) {
	return instance.SetProperty("SKU", value)
}

// GetSKU gets the value of SKU for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertySKU() (value string, err error) {
	retValue, err := instance.GetProperty("SKU")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSKUId sets the value of SKUId for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertySKUId(value uint32) (err error) {
	return instance.SetProperty("SKUId", value)
}

// GetSKUId gets the value of SKUId for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertySKUId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SKUId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSPMajorVersion sets the value of SPMajorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertySPMajorVersion(value uint16) (err error) {
	return instance.SetProperty("SPMajorVersion", value)
}

// GetSPMajorVersion gets the value of SPMajorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertySPMajorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("SPMajorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSPMinorVersion sets the value of SPMinorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) SetPropertySPMinorVersion(value uint16) (err error) {
	return instance.SetProperty("SPMinorVersion", value)
}

// GetSPMinorVersion gets the value of SPMinorVersion for the instance
func (instance *MSFT_ServerOperatingSystem) GetPropertySPMinorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("SPMinorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
