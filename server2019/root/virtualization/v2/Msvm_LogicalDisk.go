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

// Msvm_LogicalDisk struct
type Msvm_LogicalDisk struct {
	*CIM_LogicalDisk
}

func NewMsvm_LogicalDiskEx1(instance *cim.WmiInstance) (newInstance *Msvm_LogicalDisk, err error) {
	tmp, err := NewCIM_LogicalDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_LogicalDisk{
		CIM_LogicalDisk: tmp,
	}
	return
}

func NewMsvm_LogicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_LogicalDisk, err error) {
	tmp, err := NewCIM_LogicalDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_LogicalDisk{
		CIM_LogicalDisk: tmp,
	}
	return
}

func (instance *Msvm_LogicalDisk) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_LogicalDisk) GetRelatedDiskDrive() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_DiskDrive")
}

func (instance *Msvm_LogicalDisk) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_LogicalDisk) GetRelatedStorageAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_StorageAllocationSettingData")
}
