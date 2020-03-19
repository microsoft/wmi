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

// Msvm_EthernetSwitchExtension struct
type Msvm_EthernetSwitchExtension struct {
	*CIM_EnabledLogicalElement

	//
	CreationClassName string

	//
	ExtensionType uint8

	//
	SystemCreationClassName string

	//
	SystemName string

	//
	Vendor string

	//
	Version string
}

func NewMsvm_EthernetSwitchExtensionEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchExtension, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchExtension{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchExtensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchExtension, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchExtension{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtensionType sets the value of ExtensionType for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertyExtensionType(value uint8) (err error) {
	return instance.SetProperty("ExtensionType", value)
}

// GetExtensionType gets the value of ExtensionType for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertyExtensionType() (value uint8, err error) {
	retValue, err := instance.GetProperty("ExtensionType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendor sets the value of Vendor for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertyVendor(value string) (err error) {
	return instance.SetProperty("Vendor", value)
}

// GetVendor gets the value of Vendor for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertyVendor() (value string, err error) {
	retValue, err := instance.GetProperty("Vendor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *Msvm_EthernetSwitchExtension) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_EthernetSwitchExtension) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchExtension) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_EthernetSwitchExtension) GetRelatedInstalledEthernetSwitchExtension() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_InstalledEthernetSwitchExtension")
}
