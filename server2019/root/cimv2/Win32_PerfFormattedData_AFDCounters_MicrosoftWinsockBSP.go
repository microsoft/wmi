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

// Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP struct
type Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP struct {
	*Win32_PerfFormattedData

	//
	DroppedDatagrams uint32

	//
	DroppedDatagramsPersec uint32

	//
	RejectedConnections uint32

	//
	RejectedConnectionsPersec uint32
}

func NewWin32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSPEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetDroppedDatagrams sets the value of DroppedDatagrams for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyDroppedDatagrams(value uint32) (err error) {
	return instance.SetProperty("DroppedDatagrams", value)
}

// GetDroppedDatagrams gets the value of DroppedDatagrams for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyDroppedDatagrams() (value uint32, err error) {
	retValue, err := instance.GetProperty("DroppedDatagrams")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDroppedDatagramsPersec sets the value of DroppedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyDroppedDatagramsPersec(value uint32) (err error) {
	return instance.SetProperty("DroppedDatagramsPersec", value)
}

// GetDroppedDatagramsPersec gets the value of DroppedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyDroppedDatagramsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DroppedDatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRejectedConnections sets the value of RejectedConnections for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyRejectedConnections(value uint32) (err error) {
	return instance.SetProperty("RejectedConnections", value)
}

// GetRejectedConnections gets the value of RejectedConnections for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyRejectedConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("RejectedConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRejectedConnectionsPersec sets the value of RejectedConnectionsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyRejectedConnectionsPersec(value uint32) (err error) {
	return instance.SetProperty("RejectedConnectionsPersec", value)
}

// GetRejectedConnectionsPersec gets the value of RejectedConnectionsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyRejectedConnectionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RejectedConnectionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
