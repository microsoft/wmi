// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_RedundancySet struct
type Msvm_RedundancySet struct {
	*CIM_RedundancySet
}

func NewMsvm_RedundancySetEx1(instance *cim.WmiInstance) (newInstance *Msvm_RedundancySet, err error) {
	tmp, err := NewCIM_RedundancySetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_RedundancySet{
		CIM_RedundancySet: tmp,
	}
	return
}

func NewMsvm_RedundancySetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_RedundancySet, err error) {
	tmp, err := NewCIM_RedundancySetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_RedundancySet{
		CIM_RedundancySet: tmp,
	}
	return
}

func (instance *Msvm_RedundancySet) GetRelatedComputerSystem() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ComputerSystem")
}
