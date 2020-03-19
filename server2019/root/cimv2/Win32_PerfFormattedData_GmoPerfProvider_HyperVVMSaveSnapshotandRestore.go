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

// Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore struct
type Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore struct {
	*Win32_PerfFormattedData

	//
	OperationTime uint32

	//
	RequestsActive uint32

	//
	RequestsDispatched uint32

	//
	RequestsHighPriority uint32

	//
	RequestsProcessed uint32

	//
	ThreadsSpawned uint32
}

func NewWin32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestoreEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetOperationTime sets the value of OperationTime for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyOperationTime(value uint32) (err error) {
	return instance.SetProperty("OperationTime", value)
}

// GetOperationTime gets the value of OperationTime for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyOperationTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("OperationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestsActive sets the value of RequestsActive for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyRequestsActive(value uint32) (err error) {
	return instance.SetProperty("RequestsActive", value)
}

// GetRequestsActive gets the value of RequestsActive for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyRequestsActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestsActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestsDispatched sets the value of RequestsDispatched for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyRequestsDispatched(value uint32) (err error) {
	return instance.SetProperty("RequestsDispatched", value)
}

// GetRequestsDispatched gets the value of RequestsDispatched for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyRequestsDispatched() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestsDispatched")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestsHighPriority sets the value of RequestsHighPriority for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyRequestsHighPriority(value uint32) (err error) {
	return instance.SetProperty("RequestsHighPriority", value)
}

// GetRequestsHighPriority gets the value of RequestsHighPriority for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyRequestsHighPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestsHighPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestsProcessed sets the value of RequestsProcessed for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyRequestsProcessed(value uint32) (err error) {
	return instance.SetProperty("RequestsProcessed", value)
}

// GetRequestsProcessed gets the value of RequestsProcessed for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyRequestsProcessed() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestsProcessed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadsSpawned sets the value of ThreadsSpawned for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) SetPropertyThreadsSpawned(value uint32) (err error) {
	return instance.SetProperty("ThreadsSpawned", value)
}

// GetThreadsSpawned gets the value of ThreadsSpawned for the instance
func (instance *Win32_PerfFormattedData_GmoPerfProvider_HyperVVMSaveSnapshotandRestore) GetPropertyThreadsSpawned() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadsSpawned")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
