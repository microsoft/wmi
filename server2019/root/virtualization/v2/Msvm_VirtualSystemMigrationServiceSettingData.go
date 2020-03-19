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

// Msvm_VirtualSystemMigrationServiceSettingData struct
type Msvm_VirtualSystemMigrationServiceSettingData struct {
	*CIM_SettingData

	//
	AuthenticationType VirtualSystemMigrationServiceSettingData_AuthenticationType

	//
	EnableCompression bool

	//
	EnableSmbTransport bool

	//
	EnableVirtualSystemMigration bool

	//
	MaximumActiveStorageMigration uint32

	//
	MaximumActiveVirtualSystemMigration uint32
}

func NewMsvm_VirtualSystemMigrationServiceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationServiceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAuthenticationType sets the value of AuthenticationType for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyAuthenticationType(value VirtualSystemMigrationServiceSettingData_AuthenticationType) (err error) {
	return instance.SetProperty("AuthenticationType", value)
}

// GetAuthenticationType gets the value of AuthenticationType for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyAuthenticationType() (value VirtualSystemMigrationServiceSettingData_AuthenticationType, err error) {
	retValue, err := instance.GetProperty("AuthenticationType")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualSystemMigrationServiceSettingData_AuthenticationType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableCompression sets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyEnableCompression(value bool) (err error) {
	return instance.SetProperty("EnableCompression", value)
}

// GetEnableCompression gets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyEnableCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableCompression")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableSmbTransport sets the value of EnableSmbTransport for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyEnableSmbTransport(value bool) (err error) {
	return instance.SetProperty("EnableSmbTransport", value)
}

// GetEnableSmbTransport gets the value of EnableSmbTransport for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyEnableSmbTransport() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSmbTransport")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableVirtualSystemMigration sets the value of EnableVirtualSystemMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyEnableVirtualSystemMigration(value bool) (err error) {
	return instance.SetProperty("EnableVirtualSystemMigration", value)
}

// GetEnableVirtualSystemMigration gets the value of EnableVirtualSystemMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyEnableVirtualSystemMigration() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableVirtualSystemMigration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumActiveStorageMigration sets the value of MaximumActiveStorageMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyMaximumActiveStorageMigration(value uint32) (err error) {
	return instance.SetProperty("MaximumActiveStorageMigration", value)
}

// GetMaximumActiveStorageMigration gets the value of MaximumActiveStorageMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyMaximumActiveStorageMigration() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumActiveStorageMigration")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumActiveVirtualSystemMigration sets the value of MaximumActiveVirtualSystemMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) SetPropertyMaximumActiveVirtualSystemMigration(value uint32) (err error) {
	return instance.SetProperty("MaximumActiveVirtualSystemMigration", value)
}

// GetMaximumActiveVirtualSystemMigration gets the value of MaximumActiveVirtualSystemMigration for the instance
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetPropertyMaximumActiveVirtualSystemMigration() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumActiveVirtualSystemMigration")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetRelatedVirtualSystemMigrationNetworkSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_VirtualSystemMigrationNetworkSettingData")
}

func (instance *Msvm_VirtualSystemMigrationServiceSettingData) GetRelatedVirtualSystemMigrationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationService")
}
