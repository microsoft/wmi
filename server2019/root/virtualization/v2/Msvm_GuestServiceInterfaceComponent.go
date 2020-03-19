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

// Msvm_GuestServiceInterfaceComponent struct
type Msvm_GuestServiceInterfaceComponent struct {
	*CIM_LogicalDevice
}

func NewMsvm_GuestServiceInterfaceComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestServiceInterfaceComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_GuestServiceInterfaceComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestServiceInterfaceComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func (instance *Msvm_GuestServiceInterfaceComponent) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_GuestServiceInterfaceComponent) GetRelatedGuestFileService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_GuestFileService")
}

func (instance *Msvm_GuestServiceInterfaceComponent) GetRelatedGuestCommunicationService() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_GuestCommunicationService")
}

func (instance *Msvm_GuestServiceInterfaceComponent) GetRelatedGuestServiceInterfaceComponentSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_GuestServiceInterfaceComponentSettingData")
}
