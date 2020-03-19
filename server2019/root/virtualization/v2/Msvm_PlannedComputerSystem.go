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

// Msvm_PlannedComputerSystem struct
type Msvm_PlannedComputerSystem struct {
	*CIM_ComputerSystem

	//
	AssignedNumaNodeList []uint16

	//
	OnTimeInMilliseconds uint64

	//
	ProcessID uint32

	//
	TimeOfLastConfigurationChange string
}

func NewMsvm_PlannedComputerSystemEx1(instance *cim.WmiInstance) (newInstance *Msvm_PlannedComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PlannedComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMsvm_PlannedComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PlannedComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PlannedComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetAssignedNumaNodeList sets the value of AssignedNumaNodeList for the instance
func (instance *Msvm_PlannedComputerSystem) SetPropertyAssignedNumaNodeList(value []uint16) (err error) {
	return instance.SetProperty("AssignedNumaNodeList", value)
}

// GetAssignedNumaNodeList gets the value of AssignedNumaNodeList for the instance
func (instance *Msvm_PlannedComputerSystem) GetPropertyAssignedNumaNodeList() (value []uint16, err error) {
	retValue, err := instance.GetProperty("AssignedNumaNodeList")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnTimeInMilliseconds sets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_PlannedComputerSystem) SetPropertyOnTimeInMilliseconds(value uint64) (err error) {
	return instance.SetProperty("OnTimeInMilliseconds", value)
}

// GetOnTimeInMilliseconds gets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_PlannedComputerSystem) GetPropertyOnTimeInMilliseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("OnTimeInMilliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Msvm_PlannedComputerSystem) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", value)
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Msvm_PlannedComputerSystem) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastConfigurationChange sets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_PlannedComputerSystem) SetPropertyTimeOfLastConfigurationChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastConfigurationChange", value)
}

// GetTimeOfLastConfigurationChange gets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_PlannedComputerSystem) GetPropertyTimeOfLastConfigurationChange() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastConfigurationChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
