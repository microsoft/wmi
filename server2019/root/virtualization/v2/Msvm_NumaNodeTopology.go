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
	return instance.SetProperty("CountOfMemoryBlocks", value)
}

// GetCountOfMemoryBlocks gets the value of CountOfMemoryBlocks for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyCountOfMemoryBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("CountOfMemoryBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCountOfProcessors sets the value of CountOfProcessors for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyCountOfProcessors(value uint32) (err error) {
	return instance.SetProperty("CountOfProcessors", value)
}

// GetCountOfProcessors gets the value of CountOfProcessors for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyCountOfProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("CountOfProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalNodeNumber sets the value of PhysicalNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyPhysicalNodeNumber(value uint32) (err error) {
	return instance.SetProperty("PhysicalNodeNumber", value)
}

// GetPhysicalNodeNumber gets the value of PhysicalNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyPhysicalNodeNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalNodeNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualNodeNumber sets the value of VirtualNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyVirtualNodeNumber(value uint32) (err error) {
	return instance.SetProperty("VirtualNodeNumber", value)
}

// GetVirtualNodeNumber gets the value of VirtualNodeNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyVirtualNodeNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualNodeNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSocketNumber sets the value of VirtualSocketNumber for the instance
func (instance *Msvm_NumaNodeTopology) SetPropertyVirtualSocketNumber(value uint32) (err error) {
	return instance.SetProperty("VirtualSocketNumber", value)
}

// GetVirtualSocketNumber gets the value of VirtualSocketNumber for the instance
func (instance *Msvm_NumaNodeTopology) GetPropertyVirtualSocketNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSocketNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
