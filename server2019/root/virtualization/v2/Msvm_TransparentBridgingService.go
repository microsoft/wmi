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

// Msvm_TransparentBridgingService struct
type Msvm_TransparentBridgingService struct {
	CIM_TransparentBridgingService
}

func (instance *Msvm_TransparentBridgingService) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_TransparentBridgingService) GetRelatedDynamicForwardingEntry() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_DynamicForwardingEntry")
}
