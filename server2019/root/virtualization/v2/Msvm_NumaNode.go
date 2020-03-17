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

// Msvm_NumaNode struct
type Msvm_NumaNode struct {
	CIM_EnabledLogicalElement

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	//
	CurrentlyAssignedVirtualProcessors uint32

	//
	CurrentlyConsumableMemoryBlocks uint64

	//
	NodeID string

	//
	NumberOfLogicalProcessors uint32

	//
	NumberOfProcessorCores uint32

	// The scoping System's CreationClassName.
	SystemCreationClassName string

	// The scoping System's Name.
	SystemName string
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *Msvm_NumaNode) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *Msvm_NumaNode) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentlyAssignedVirtualProcessors sets the value of CurrentlyAssignedVirtualProcessors for the instance
func (instance *Msvm_NumaNode) SetPropertyCurrentlyAssignedVirtualProcessors(value uint32) (err error) {
	return instance.SetProperty("CurrentlyAssignedVirtualProcessors", value)
}

// GetCurrentlyAssignedVirtualProcessors gets the value of CurrentlyAssignedVirtualProcessors for the instance
func (instance *Msvm_NumaNode) GetPropertyCurrentlyAssignedVirtualProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentlyAssignedVirtualProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentlyConsumableMemoryBlocks sets the value of CurrentlyConsumableMemoryBlocks for the instance
func (instance *Msvm_NumaNode) SetPropertyCurrentlyConsumableMemoryBlocks(value uint64) (err error) {
	return instance.SetProperty("CurrentlyConsumableMemoryBlocks", value)
}

// GetCurrentlyConsumableMemoryBlocks gets the value of CurrentlyConsumableMemoryBlocks for the instance
func (instance *Msvm_NumaNode) GetPropertyCurrentlyConsumableMemoryBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentlyConsumableMemoryBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeID sets the value of NodeID for the instance
func (instance *Msvm_NumaNode) SetPropertyNodeID(value string) (err error) {
	return instance.SetProperty("NodeID", value)
}

// GetNodeID gets the value of NodeID for the instance
func (instance *Msvm_NumaNode) GetPropertyNodeID() (value string, err error) {
	retValue, err := instance.GetProperty("NodeID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfLogicalProcessors sets the value of NumberOfLogicalProcessors for the instance
func (instance *Msvm_NumaNode) SetPropertyNumberOfLogicalProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfLogicalProcessors", value)
}

// GetNumberOfLogicalProcessors gets the value of NumberOfLogicalProcessors for the instance
func (instance *Msvm_NumaNode) GetPropertyNumberOfLogicalProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfLogicalProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfProcessorCores sets the value of NumberOfProcessorCores for the instance
func (instance *Msvm_NumaNode) SetPropertyNumberOfProcessorCores(value uint32) (err error) {
	return instance.SetProperty("NumberOfProcessorCores", value)
}

// GetNumberOfProcessorCores gets the value of NumberOfProcessorCores for the instance
func (instance *Msvm_NumaNode) GetPropertyNumberOfProcessorCores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessorCores")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *Msvm_NumaNode) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *Msvm_NumaNode) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *Msvm_NumaNode) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *Msvm_NumaNode) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_NumaNode) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_NumaNode) GetRelatedMemory() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_Memory")
}

func (instance *Msvm_NumaNode) GetRelatedProcessor() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Processor")
}
