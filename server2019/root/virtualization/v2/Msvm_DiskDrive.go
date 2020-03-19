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

// Msvm_DiskDrive struct
type Msvm_DiskDrive struct {
	*CIM_DiskDrive

	//
	DriveNumber uint32
}

func NewMsvm_DiskDriveEx1(instance *cim.WmiInstance) (newInstance *Msvm_DiskDrive, err error) {
	tmp, err := NewCIM_DiskDriveEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_DiskDrive{
		CIM_DiskDrive: tmp,
	}
	return
}

func NewMsvm_DiskDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_DiskDrive, err error) {
	tmp, err := NewCIM_DiskDriveEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_DiskDrive{
		CIM_DiskDrive: tmp,
	}
	return
}

// SetDriveNumber sets the value of DriveNumber for the instance
func (instance *Msvm_DiskDrive) SetPropertyDriveNumber(value uint32) (err error) {
	return instance.SetProperty("DriveNumber", value)
}

// GetDriveNumber gets the value of DriveNumber for the instance
func (instance *Msvm_DiskDrive) GetPropertyDriveNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_DiskDrive) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_DiskDrive) GetRelatedLogicalDisk() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LogicalDisk")
}

func (instance *Msvm_DiskDrive) GetRelatedResourceAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourceAllocationSettingData")
}

func (instance *Msvm_DiskDrive) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_DiskDrive) GetRelatedSCSIProtocolController() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SCSIProtocolController")
}
