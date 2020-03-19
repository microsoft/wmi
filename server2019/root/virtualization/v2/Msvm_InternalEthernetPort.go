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

// Msvm_InternalEthernetPort struct
type Msvm_InternalEthernetPort struct {
	*CIM_EthernetPort
}

func NewMsvm_InternalEthernetPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_InternalEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_InternalEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func NewMsvm_InternalEthernetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_InternalEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_InternalEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func (instance *Msvm_InternalEthernetPort) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_InternalEthernetPort) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}
