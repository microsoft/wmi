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

// Msvm_SyntheticEthernetPort struct
type Msvm_SyntheticEthernetPort struct {
	*CIM_EthernetPort
}

func NewMsvm_SyntheticEthernetPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func NewMsvm_SyntheticEthernetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func (instance *Msvm_SyntheticEthernetPort) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_SyntheticEthernetPort) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_SyntheticEthernetPort) GetRelatedSyntheticEthernetPortSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SyntheticEthernetPortSettingData")
}

func (instance *Msvm_SyntheticEthernetPort) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
