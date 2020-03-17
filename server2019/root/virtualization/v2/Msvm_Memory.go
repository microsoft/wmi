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

// Msvm_Memory struct
type Msvm_Memory struct {
	CIM_Memory
}

func (instance *Msvm_Memory) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_Memory) GetRelatedMemory() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Memory")
}
