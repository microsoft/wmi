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

// Msvm_InstalledEthernetSwitchExtension struct
type Msvm_InstalledEthernetSwitchExtension struct {
	*CIM_ManagedSystemElement

	//
	ExtensionType uint8

	//
	Vendor string

	//
	Version string
}

func NewMsvm_InstalledEthernetSwitchExtensionEx1(instance *cim.WmiInstance) (newInstance *Msvm_InstalledEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_InstalledEthernetSwitchExtension{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

func NewMsvm_InstalledEthernetSwitchExtensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_InstalledEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_InstalledEthernetSwitchExtension{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

// SetExtensionType sets the value of ExtensionType for the instance
func (instance *Msvm_InstalledEthernetSwitchExtension) SetPropertyExtensionType(value uint8) (err error) {
	return instance.SetProperty("ExtensionType", value)
}

// GetExtensionType gets the value of ExtensionType for the instance
func (instance *Msvm_InstalledEthernetSwitchExtension) GetPropertyExtensionType() (value uint8, err error) {
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

// SetVendor sets the value of Vendor for the instance
func (instance *Msvm_InstalledEthernetSwitchExtension) SetPropertyVendor(value string) (err error) {
	return instance.SetProperty("Vendor", value)
}

// GetVendor gets the value of Vendor for the instance
func (instance *Msvm_InstalledEthernetSwitchExtension) GetPropertyVendor() (value string, err error) {
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
func (instance *Msvm_InstalledEthernetSwitchExtension) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_InstalledEthernetSwitchExtension) GetPropertyVersion() (value string, err error) {
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
func (instance *Msvm_InstalledEthernetSwitchExtension) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_InstalledEthernetSwitchExtension) GetRelatedEthernetSwitchExtension() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchExtension")
}
