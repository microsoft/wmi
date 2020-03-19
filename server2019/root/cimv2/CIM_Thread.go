// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Thread struct
type CIM_Thread struct {
	*CIM_LogicalElement

	//
	CreationClassName string

	//
	CSCreationClassName string

	//
	CSName string

	//
	ExecutionState uint16

	//
	Handle string

	//
	KernelModeTime uint64

	//
	OSCreationClassName string

	//
	OSName string

	//
	Priority uint32

	//
	ProcessCreationClassName string

	//
	ProcessHandle string

	//
	UserModeTime uint64
}

func NewCIM_ThreadEx1(instance *cim.WmiInstance) (newInstance *CIM_Thread, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Thread{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_ThreadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Thread, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Thread{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Thread) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Thread) GetPropertyCreationClassName() (value string, err error) {
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

// SetCSCreationClassName sets the value of CSCreationClassName for the instance
func (instance *CIM_Thread) SetPropertyCSCreationClassName(value string) (err error) {
	return instance.SetProperty("CSCreationClassName", value)
}

// GetCSCreationClassName gets the value of CSCreationClassName for the instance
func (instance *CIM_Thread) GetPropertyCSCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CSCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCSName sets the value of CSName for the instance
func (instance *CIM_Thread) SetPropertyCSName(value string) (err error) {
	return instance.SetProperty("CSName", value)
}

// GetCSName gets the value of CSName for the instance
func (instance *CIM_Thread) GetPropertyCSName() (value string, err error) {
	retValue, err := instance.GetProperty("CSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExecutionState sets the value of ExecutionState for the instance
func (instance *CIM_Thread) SetPropertyExecutionState(value uint16) (err error) {
	return instance.SetProperty("ExecutionState", value)
}

// GetExecutionState gets the value of ExecutionState for the instance
func (instance *CIM_Thread) GetPropertyExecutionState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExecutionState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHandle sets the value of Handle for the instance
func (instance *CIM_Thread) SetPropertyHandle(value string) (err error) {
	return instance.SetProperty("Handle", value)
}

// GetHandle gets the value of Handle for the instance
func (instance *CIM_Thread) GetPropertyHandle() (value string, err error) {
	retValue, err := instance.GetProperty("Handle")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKernelModeTime sets the value of KernelModeTime for the instance
func (instance *CIM_Thread) SetPropertyKernelModeTime(value uint64) (err error) {
	return instance.SetProperty("KernelModeTime", value)
}

// GetKernelModeTime gets the value of KernelModeTime for the instance
func (instance *CIM_Thread) GetPropertyKernelModeTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("KernelModeTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSCreationClassName sets the value of OSCreationClassName for the instance
func (instance *CIM_Thread) SetPropertyOSCreationClassName(value string) (err error) {
	return instance.SetProperty("OSCreationClassName", value)
}

// GetOSCreationClassName gets the value of OSCreationClassName for the instance
func (instance *CIM_Thread) GetPropertyOSCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("OSCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSName sets the value of OSName for the instance
func (instance *CIM_Thread) SetPropertyOSName(value string) (err error) {
	return instance.SetProperty("OSName", value)
}

// GetOSName gets the value of OSName for the instance
func (instance *CIM_Thread) GetPropertyOSName() (value string, err error) {
	retValue, err := instance.GetProperty("OSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *CIM_Thread) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_Thread) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessCreationClassName sets the value of ProcessCreationClassName for the instance
func (instance *CIM_Thread) SetPropertyProcessCreationClassName(value string) (err error) {
	return instance.SetProperty("ProcessCreationClassName", value)
}

// GetProcessCreationClassName gets the value of ProcessCreationClassName for the instance
func (instance *CIM_Thread) GetPropertyProcessCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessHandle sets the value of ProcessHandle for the instance
func (instance *CIM_Thread) SetPropertyProcessHandle(value string) (err error) {
	return instance.SetProperty("ProcessHandle", value)
}

// GetProcessHandle gets the value of ProcessHandle for the instance
func (instance *CIM_Thread) GetPropertyProcessHandle() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessHandle")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserModeTime sets the value of UserModeTime for the instance
func (instance *CIM_Thread) SetPropertyUserModeTime(value uint64) (err error) {
	return instance.SetProperty("UserModeTime", value)
}

// GetUserModeTime gets the value of UserModeTime for the instance
func (instance *CIM_Thread) GetPropertyUserModeTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("UserModeTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
