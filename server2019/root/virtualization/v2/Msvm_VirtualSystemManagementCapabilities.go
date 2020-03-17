// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemManagementCapabilities struct
type Msvm_VirtualSystemManagementCapabilities struct {
	CIM_VirtualSystemManagementCapabilities
}

func (instance *Msvm_VirtualSystemManagementCapabilities) GetRelatedVirtualSystemManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemManagementService")
}

func (instance *Msvm_VirtualSystemManagementCapabilities) GetRelatedVirtualSystemSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_VirtualSystemSettingData")
}
