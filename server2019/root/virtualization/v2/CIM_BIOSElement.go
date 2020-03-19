// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BIOSElement struct
type CIM_BIOSElement struct {
	*CIM_SoftwareElement

	// The currently selected language for the BIOS. This information can be obtained from SMBIOS, using the Current Language attribute of the Type 13 structure, to index into the string list following the structure. The property is formatted using the ISO 639 Language Name, and may be followed by the ISO 3166 Territory Name and the encoding method.
	CurrentLanguage string

	// A list of installable languages for the BIOS. This information can be obtained from SMBIOS, from the string list that follows the Type 13 structure. An ISO 639 Language Name should be used to specify the BIOS' installable languages. The ISO 3166 Territory Name and the encoding method may also be specified, following the Language Name.
	ListOfLanguages []string

	// The ending address of the memory which this BIOS occupies.
	LoadedEndingAddress uint64

	// The starting address of the memory which this BIOS occupies.
	LoadedStartingAddress uint64

	// A free form string describing the BIOS flash/load utility that is required to update the BIOSElement. Version and other information may be indicated in this property.
	LoadUtilityInformation string

	// If true, this is the primary BIOS of the ComputerSystem.
	PrimaryBIOS bool

	// A string representing the publication location of the BIOS Attribute registry or registries the implementation complies to.
	RegistryURIs []string

	// Date that this BIOS was released.
	ReleaseDate string
}

func NewCIM_BIOSElementEx1(instance *cim.WmiInstance) (newInstance *CIM_BIOSElement, err error) {
	tmp, err := NewCIM_SoftwareElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSElement{
		CIM_SoftwareElement: tmp,
	}
	return
}

func NewCIM_BIOSElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BIOSElement, err error) {
	tmp, err := NewCIM_SoftwareElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSElement{
		CIM_SoftwareElement: tmp,
	}
	return
}

// SetCurrentLanguage sets the value of CurrentLanguage for the instance
func (instance *CIM_BIOSElement) SetPropertyCurrentLanguage(value string) (err error) {
	return instance.SetProperty("CurrentLanguage", value)
}

// GetCurrentLanguage gets the value of CurrentLanguage for the instance
func (instance *CIM_BIOSElement) GetPropertyCurrentLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentLanguage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetListOfLanguages sets the value of ListOfLanguages for the instance
func (instance *CIM_BIOSElement) SetPropertyListOfLanguages(value []string) (err error) {
	return instance.SetProperty("ListOfLanguages", value)
}

// GetListOfLanguages gets the value of ListOfLanguages for the instance
func (instance *CIM_BIOSElement) GetPropertyListOfLanguages() (value []string, err error) {
	retValue, err := instance.GetProperty("ListOfLanguages")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadedEndingAddress sets the value of LoadedEndingAddress for the instance
func (instance *CIM_BIOSElement) SetPropertyLoadedEndingAddress(value uint64) (err error) {
	return instance.SetProperty("LoadedEndingAddress", value)
}

// GetLoadedEndingAddress gets the value of LoadedEndingAddress for the instance
func (instance *CIM_BIOSElement) GetPropertyLoadedEndingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("LoadedEndingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadedStartingAddress sets the value of LoadedStartingAddress for the instance
func (instance *CIM_BIOSElement) SetPropertyLoadedStartingAddress(value uint64) (err error) {
	return instance.SetProperty("LoadedStartingAddress", value)
}

// GetLoadedStartingAddress gets the value of LoadedStartingAddress for the instance
func (instance *CIM_BIOSElement) GetPropertyLoadedStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("LoadedStartingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadUtilityInformation sets the value of LoadUtilityInformation for the instance
func (instance *CIM_BIOSElement) SetPropertyLoadUtilityInformation(value string) (err error) {
	return instance.SetProperty("LoadUtilityInformation", value)
}

// GetLoadUtilityInformation gets the value of LoadUtilityInformation for the instance
func (instance *CIM_BIOSElement) GetPropertyLoadUtilityInformation() (value string, err error) {
	retValue, err := instance.GetProperty("LoadUtilityInformation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryBIOS sets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) SetPropertyPrimaryBIOS(value bool) (err error) {
	return instance.SetProperty("PrimaryBIOS", value)
}

// GetPrimaryBIOS gets the value of PrimaryBIOS for the instance
func (instance *CIM_BIOSElement) GetPropertyPrimaryBIOS() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryBIOS")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryURIs sets the value of RegistryURIs for the instance
func (instance *CIM_BIOSElement) SetPropertyRegistryURIs(value []string) (err error) {
	return instance.SetProperty("RegistryURIs", value)
}

// GetRegistryURIs gets the value of RegistryURIs for the instance
func (instance *CIM_BIOSElement) GetPropertyRegistryURIs() (value []string, err error) {
	retValue, err := instance.GetProperty("RegistryURIs")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReleaseDate sets the value of ReleaseDate for the instance
func (instance *CIM_BIOSElement) SetPropertyReleaseDate(value string) (err error) {
	return instance.SetProperty("ReleaseDate", value)
}

// GetReleaseDate gets the value of ReleaseDate for the instance
func (instance *CIM_BIOSElement) GetPropertyReleaseDate() (value string, err error) {
	retValue, err := instance.GetProperty("ReleaseDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
