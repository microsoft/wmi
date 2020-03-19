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

// Msvm_TransparentBridgingService struct
type Msvm_TransparentBridgingService struct {
	*CIM_TransparentBridgingService
}

func NewMsvm_TransparentBridgingServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_TransparentBridgingService, err error) {
	tmp, err := NewCIM_TransparentBridgingServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TransparentBridgingService{
		CIM_TransparentBridgingService: tmp,
	}
	return
}

func NewMsvm_TransparentBridgingServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TransparentBridgingService, err error) {
	tmp, err := NewCIM_TransparentBridgingServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TransparentBridgingService{
		CIM_TransparentBridgingService: tmp,
	}
	return
}

func (instance *Msvm_TransparentBridgingService) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_TransparentBridgingService) GetRelatedDynamicForwardingEntry() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_DynamicForwardingEntry")
}
