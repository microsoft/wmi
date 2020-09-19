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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_DVDDrive struct
type Msvm_DVDDrive struct {
	*CIM_DVDDrive
}

func NewMsvm_DVDDriveEx1(instance *cim.WmiInstance) (newInstance *Msvm_DVDDrive, err error) {
	tmp, err := NewCIM_DVDDriveEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_DVDDrive{
		CIM_DVDDrive: tmp,
	}
	return
}

func NewMsvm_DVDDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_DVDDrive, err error) {
	tmp, err := NewCIM_DVDDriveEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_DVDDrive{
		CIM_DVDDrive: tmp,
	}
	return
}

func (instance *Msvm_DVDDrive) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_DVDDrive) GetRelatedLogicalDisk() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LogicalDisk")
}

func (instance *Msvm_DVDDrive) GetRelatedResourceAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourceAllocationSettingData")
}

func (instance *Msvm_DVDDrive) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_DVDDrive) GetRelatedSCSIProtocolController() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SCSIProtocolController")
}
