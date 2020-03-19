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

// Win32_NamedJobObjectActgInfo struct
type Win32_NamedJobObjectActgInfo struct {
	*CIM_StatisticalInformation

	//
	ActiveProcesses uint32

	//
	OtherOperationCount uint64

	//
	OtherTransferCount uint64

	//
	PeakJobMemoryUsed uint32

	//
	PeakProcessMemoryUsed uint32

	//
	ReadOperationCount uint64

	//
	ReadTransferCount uint64

	//
	ThisPeriodTotalKernelTime uint64

	//
	ThisPeriodTotalUserTime uint64

	//
	TotalKernelTime uint64

	//
	TotalPageFaultCount uint32

	//
	TotalProcesses uint32

	//
	TotalTerminatedProcesses uint32

	//
	TotalUserTime uint64

	//
	WriteOperationCount uint64

	//
	WriteTransferCount uint64
}

func NewWin32_NamedJobObjectActgInfoEx1(instance *cim.WmiInstance) (newInstance *Win32_NamedJobObjectActgInfo, err error) {
	tmp, err := NewCIM_StatisticalInformationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectActgInfo{
		CIM_StatisticalInformation: tmp,
	}
	return
}

func NewWin32_NamedJobObjectActgInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NamedJobObjectActgInfo, err error) {
	tmp, err := NewCIM_StatisticalInformationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectActgInfo{
		CIM_StatisticalInformation: tmp,
	}
	return
}

// SetActiveProcesses sets the value of ActiveProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyActiveProcesses(value uint32) (err error) {
	return instance.SetProperty("ActiveProcesses", value)
}

// GetActiveProcesses gets the value of ActiveProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyActiveProcesses() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherOperationCount sets the value of OtherOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyOtherOperationCount(value uint64) (err error) {
	return instance.SetProperty("OtherOperationCount", value)
}

// GetOtherOperationCount gets the value of OtherOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyOtherOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherOperationCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTransferCount sets the value of OtherTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyOtherTransferCount(value uint64) (err error) {
	return instance.SetProperty("OtherTransferCount", value)
}

// GetOtherTransferCount gets the value of OtherTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyOtherTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherTransferCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeakJobMemoryUsed sets the value of PeakJobMemoryUsed for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyPeakJobMemoryUsed(value uint32) (err error) {
	return instance.SetProperty("PeakJobMemoryUsed", value)
}

// GetPeakJobMemoryUsed gets the value of PeakJobMemoryUsed for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyPeakJobMemoryUsed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeakJobMemoryUsed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeakProcessMemoryUsed sets the value of PeakProcessMemoryUsed for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyPeakProcessMemoryUsed(value uint32) (err error) {
	return instance.SetProperty("PeakProcessMemoryUsed", value)
}

// GetPeakProcessMemoryUsed gets the value of PeakProcessMemoryUsed for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyPeakProcessMemoryUsed() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeakProcessMemoryUsed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadOperationCount sets the value of ReadOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyReadOperationCount(value uint64) (err error) {
	return instance.SetProperty("ReadOperationCount", value)
}

// GetReadOperationCount gets the value of ReadOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyReadOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadOperationCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadTransferCount sets the value of ReadTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyReadTransferCount(value uint64) (err error) {
	return instance.SetProperty("ReadTransferCount", value)
}

// GetReadTransferCount gets the value of ReadTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyReadTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadTransferCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThisPeriodTotalKernelTime sets the value of ThisPeriodTotalKernelTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyThisPeriodTotalKernelTime(value uint64) (err error) {
	return instance.SetProperty("ThisPeriodTotalKernelTime", value)
}

// GetThisPeriodTotalKernelTime gets the value of ThisPeriodTotalKernelTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyThisPeriodTotalKernelTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThisPeriodTotalKernelTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThisPeriodTotalUserTime sets the value of ThisPeriodTotalUserTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyThisPeriodTotalUserTime(value uint64) (err error) {
	return instance.SetProperty("ThisPeriodTotalUserTime", value)
}

// GetThisPeriodTotalUserTime gets the value of ThisPeriodTotalUserTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyThisPeriodTotalUserTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThisPeriodTotalUserTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalKernelTime sets the value of TotalKernelTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyTotalKernelTime(value uint64) (err error) {
	return instance.SetProperty("TotalKernelTime", value)
}

// GetTotalKernelTime gets the value of TotalKernelTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyTotalKernelTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalKernelTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPageFaultCount sets the value of TotalPageFaultCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyTotalPageFaultCount(value uint32) (err error) {
	return instance.SetProperty("TotalPageFaultCount", value)
}

// GetTotalPageFaultCount gets the value of TotalPageFaultCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyTotalPageFaultCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalPageFaultCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalProcesses sets the value of TotalProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyTotalProcesses(value uint32) (err error) {
	return instance.SetProperty("TotalProcesses", value)
}

// GetTotalProcesses gets the value of TotalProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyTotalProcesses() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalTerminatedProcesses sets the value of TotalTerminatedProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyTotalTerminatedProcesses(value uint32) (err error) {
	return instance.SetProperty("TotalTerminatedProcesses", value)
}

// GetTotalTerminatedProcesses gets the value of TotalTerminatedProcesses for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyTotalTerminatedProcesses() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalTerminatedProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalUserTime sets the value of TotalUserTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyTotalUserTime(value uint64) (err error) {
	return instance.SetProperty("TotalUserTime", value)
}

// GetTotalUserTime gets the value of TotalUserTime for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyTotalUserTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalUserTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteOperationCount sets the value of WriteOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyWriteOperationCount(value uint64) (err error) {
	return instance.SetProperty("WriteOperationCount", value)
}

// GetWriteOperationCount gets the value of WriteOperationCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyWriteOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteOperationCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteTransferCount sets the value of WriteTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) SetPropertyWriteTransferCount(value uint64) (err error) {
	return instance.SetProperty("WriteTransferCount", value)
}

// GetWriteTransferCount gets the value of WriteTransferCount for the instance
func (instance *Win32_NamedJobObjectActgInfo) GetPropertyWriteTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteTransferCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
