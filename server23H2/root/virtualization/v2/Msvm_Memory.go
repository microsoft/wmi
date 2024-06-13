// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_Memory struct
type Msvm_Memory struct {
	*CIM_Memory

	//
	MemoryEncryption bool
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

// SetMemoryEncryption sets the value of MemoryEncryption for the instance
func (instance *Msvm_Memory) SetPropertyMemoryEncryption(value bool) (err error) {
	return instance.SetProperty("MemoryEncryption", (value))
}

// GetMemoryEncryption gets the value of MemoryEncryption for the instance
func (instance *Msvm_Memory) GetPropertyMemoryEncryption() (value bool, err error) {
	retValue, err := instance.GetProperty("MemoryEncryption")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
func (instance *Msvm_Memory) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_Memory) GetRelatedMemory() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Memory")
}
