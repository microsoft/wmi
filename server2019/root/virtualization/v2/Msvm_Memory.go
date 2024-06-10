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

// Msvm_Memory struct
type Msvm_Memory struct {
	*CIM_Memory
}

func NewMsvm_MemoryEx1(instance *cim.WmiInstance) (newInstance *Msvm_Memory, err error) {
	tmp, err := NewCIM_MemoryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Memory{
		CIM_Memory: tmp,
	}
	return
}

func NewMsvm_MemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Memory, err error) {
	tmp, err := NewCIM_MemoryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Memory{
		CIM_Memory: tmp,
	}
	return
}

func (instance *Msvm_Memory) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_Memory) GetRelatedMemory() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Memory")
}
