// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_NumaNode struct
type Msvm_NumaNode struct {
	*CIM_EnabledLogicalElement

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

func NewMsvm_NumaNodeEx1(instance *cim.WmiInstance) (newInstance *Msvm_NumaNode, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_NumaNode{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewMsvm_NumaNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_NumaNode, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_NumaNode{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *Msvm_NumaNode) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *Msvm_NumaNode) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCurrentlyAssignedVirtualProcessors sets the value of CurrentlyAssignedVirtualProcessors for the instance
func (instance *Msvm_NumaNode) SetPropertyCurrentlyAssignedVirtualProcessors(value uint32) (err error) {
	return instance.SetProperty("CurrentlyAssignedVirtualProcessors", (value))
}

// GetCurrentlyAssignedVirtualProcessors gets the value of CurrentlyAssignedVirtualProcessors for the instance
func (instance *Msvm_NumaNode) GetPropertyCurrentlyAssignedVirtualProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentlyAssignedVirtualProcessors")
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

// SetCurrentlyConsumableMemoryBlocks sets the value of CurrentlyConsumableMemoryBlocks for the instance
func (instance *Msvm_NumaNode) SetPropertyCurrentlyConsumableMemoryBlocks(value uint64) (err error) {
	return instance.SetProperty("CurrentlyConsumableMemoryBlocks", (value))
}

// GetCurrentlyConsumableMemoryBlocks gets the value of CurrentlyConsumableMemoryBlocks for the instance
func (instance *Msvm_NumaNode) GetPropertyCurrentlyConsumableMemoryBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentlyConsumableMemoryBlocks")
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

// SetNodeID sets the value of NodeID for the instance
func (instance *Msvm_NumaNode) SetPropertyNodeID(value string) (err error) {
	return instance.SetProperty("NodeID", (value))
}

// GetNodeID gets the value of NodeID for the instance
func (instance *Msvm_NumaNode) GetPropertyNodeID() (value string, err error) {
	retValue, err := instance.GetProperty("NodeID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNumberOfLogicalProcessors sets the value of NumberOfLogicalProcessors for the instance
func (instance *Msvm_NumaNode) SetPropertyNumberOfLogicalProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfLogicalProcessors", (value))
}

// GetNumberOfLogicalProcessors gets the value of NumberOfLogicalProcessors for the instance
func (instance *Msvm_NumaNode) GetPropertyNumberOfLogicalProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfLogicalProcessors")
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

// SetNumberOfProcessorCores sets the value of NumberOfProcessorCores for the instance
func (instance *Msvm_NumaNode) SetPropertyNumberOfProcessorCores(value uint32) (err error) {
	return instance.SetProperty("NumberOfProcessorCores", (value))
}

// GetNumberOfProcessorCores gets the value of NumberOfProcessorCores for the instance
func (instance *Msvm_NumaNode) GetPropertyNumberOfProcessorCores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessorCores")
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

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *Msvm_NumaNode) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *Msvm_NumaNode) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *Msvm_NumaNode) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *Msvm_NumaNode) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

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
