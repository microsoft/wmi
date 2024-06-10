// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemManagementCapabilities struct
type Msvm_VirtualSystemManagementCapabilities struct {
	*CIM_VirtualSystemManagementCapabilities
}

func NewMsvm_VirtualSystemManagementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemManagementCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemManagementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemManagementCapabilities{
		CIM_VirtualSystemManagementCapabilities: tmp,
	}
	return
}

func NewMsvm_VirtualSystemManagementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemManagementCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemManagementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemManagementCapabilities{
		CIM_VirtualSystemManagementCapabilities: tmp,
	}
	return
}

func (instance *Msvm_VirtualSystemManagementCapabilities) GetRelatedVirtualSystemManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemManagementService")
}

func (instance *Msvm_VirtualSystemManagementCapabilities) GetRelatedVirtualSystemSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_VirtualSystemSettingData")
}
