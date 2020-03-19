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

// Msvm_SecurityElement struct
type Msvm_SecurityElement struct {
	*CIM_EnabledLogicalElement

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	//
	EncryptStateAndVmMigrationTrafficEnabled bool

	//
	Shielded bool

	// The scoping System's CreationClassName.
	SystemCreationClassName string

	// The scoping System's Name.
	SystemName string
}

func NewMsvm_SecurityElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_SecurityElement, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SecurityElement{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewMsvm_SecurityElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SecurityElement, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SecurityElement{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *Msvm_SecurityElement) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *Msvm_SecurityElement) GetPropertyCreationClassName() (value string, err error) {
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

// SetEncryptStateAndVmMigrationTrafficEnabled sets the value of EncryptStateAndVmMigrationTrafficEnabled for the instance
func (instance *Msvm_SecurityElement) SetPropertyEncryptStateAndVmMigrationTrafficEnabled(value bool) (err error) {
	return instance.SetProperty("EncryptStateAndVmMigrationTrafficEnabled", value)
}

// GetEncryptStateAndVmMigrationTrafficEnabled gets the value of EncryptStateAndVmMigrationTrafficEnabled for the instance
func (instance *Msvm_SecurityElement) GetPropertyEncryptStateAndVmMigrationTrafficEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptStateAndVmMigrationTrafficEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShielded sets the value of Shielded for the instance
func (instance *Msvm_SecurityElement) SetPropertyShielded(value bool) (err error) {
	return instance.SetProperty("Shielded", value)
}

// GetShielded gets the value of Shielded for the instance
func (instance *Msvm_SecurityElement) GetPropertyShielded() (value bool, err error) {
	retValue, err := instance.GetProperty("Shielded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *Msvm_SecurityElement) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *Msvm_SecurityElement) GetPropertySystemCreationClassName() (value string, err error) {
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
func (instance *Msvm_SecurityElement) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *Msvm_SecurityElement) GetPropertySystemName() (value string, err error) {
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
func (instance *Msvm_SecurityElement) GetRelatedSecuritySettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SecuritySettingData")
}

func (instance *Msvm_SecurityElement) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
