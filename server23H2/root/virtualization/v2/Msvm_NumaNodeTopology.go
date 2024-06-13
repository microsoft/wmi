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

// Msvm_NumaNodeTopology struct
type Msvm_NumaNodeTopology struct {
	*CIM_ManagedElement

	//
	CountOfMemoryBlocks uint64

	//
	CountOfProcessors uint32

	//
	PhysicalNodeNumber uint32

	//
	VirtualNodeNumber uint32

	//
	VirtualSocketNumber uint32
}

func NewMsvm_NumaNodeTopologyEx1(instance *cim.WmiInstance) (newInstance *Msvm_NumaNodeTopology, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_NumaNodeTopology{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_NumaNodeTopologyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_NumaNodeTopology, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_NumaNodeTopology{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetCountOfMemoryBlocks sets the value of CountOfMemoryBlocks for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyCountOfMemoryBlocks(value uint64) (err error) {
	return instance.SetProperty("CountOfMemoryBlocks", (value))
}

// GetCountOfMemoryBlocks gets the value of CountOfMemoryBlocks for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyCountOfMemoryBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("CountOfMemoryBlocks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCountOfProcessors sets the value of CountOfProcessors for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyCountOfProcessors(value uint32) (err error) {
	return instance.SetProperty("CountOfProcessors", (value))
}

// GetCountOfProcessors gets the value of CountOfProcessors for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyCountOfProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("CountOfProcessors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPhysicalNodeNumber sets the value of PhysicalNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyPhysicalNodeNumber(value uint32) (err error) {
	return instance.SetProperty("PhysicalNodeNumber", (value))
}

// GetPhysicalNodeNumber gets the value of PhysicalNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyPhysicalNodeNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalNodeNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVirtualNodeNumber sets the value of VirtualNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyVirtualNodeNumber(value uint32) (err error) {
	return instance.SetProperty("VirtualNodeNumber", (value))
}

// GetVirtualNodeNumber gets the value of VirtualNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyVirtualNodeNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualNodeNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVirtualSocketNumber sets the value of VirtualSocketNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyVirtualSocketNumber(value uint32) (err error) {
	return instance.SetProperty("VirtualSocketNumber", (value))
}

// GetVirtualSocketNumber gets the value of VirtualSocketNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyVirtualSocketNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSocketNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
