// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterVPortSettingData struct
type MSFT_NetAdapterVPortSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	FilterList []MSFT_NetAdapter_VmqFilter

	//
	FunctionID uint16

	//
	InterruptModeration uint32

	//
	NumFilters uint32

	//
	NumQueuePairs uint32

	//
	ProcessorAffinityMask uint64

	//
	ProcessorGroup uint16

	//
	SwitchID uint32

	//
	VPortID uint32

	//
	VPortName string

	//
	VPortState uint32
}

func NewMSFT_NetAdapterVPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterVPortSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVPortSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterVPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterVPortSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVPortSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetFilterList sets the value of FilterList for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyFilterList(value []MSFT_NetAdapter_VmqFilter) (err error) {
	return instance.SetProperty("FilterList", value)
}

// GetFilterList gets the value of FilterList for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyFilterList() (value []MSFT_NetAdapter_VmqFilter, err error) {
	retValue, err := instance.GetProperty("FilterList")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_NetAdapter_VmqFilter)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionID sets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyFunctionID(value uint16) (err error) {
	return instance.SetProperty("FunctionID", value)
}

// GetFunctionID gets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyFunctionID() (value uint16, err error) {
	retValue, err := instance.GetProperty("FunctionID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterruptModeration sets the value of InterruptModeration for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyInterruptModeration(value uint32) (err error) {
	return instance.SetProperty("InterruptModeration", value)
}

// GetInterruptModeration gets the value of InterruptModeration for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyInterruptModeration() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterruptModeration")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumFilters sets the value of NumFilters for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyNumFilters(value uint32) (err error) {
	return instance.SetProperty("NumFilters", value)
}

// GetNumFilters gets the value of NumFilters for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyNumFilters() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumFilters")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumQueuePairs sets the value of NumQueuePairs for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyNumQueuePairs(value uint32) (err error) {
	return instance.SetProperty("NumQueuePairs", value)
}

// GetNumQueuePairs gets the value of NumQueuePairs for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyNumQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorAffinityMask sets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyProcessorAffinityMask(value uint64) (err error) {
	return instance.SetProperty("ProcessorAffinityMask", value)
}

// GetProcessorAffinityMask gets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyProcessorAffinityMask() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessorAffinityMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorGroup sets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", value)
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchID sets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertySwitchID(value uint32) (err error) {
	return instance.SetProperty("SwitchID", value)
}

// GetSwitchID gets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertySwitchID() (value uint32, err error) {
	retValue, err := instance.GetProperty("SwitchID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVPortID sets the value of VPortID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyVPortID(value uint32) (err error) {
	return instance.SetProperty("VPortID", value)
}

// GetVPortID gets the value of VPortID for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyVPortID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VPortID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVPortName sets the value of VPortName for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyVPortName(value string) (err error) {
	return instance.SetProperty("VPortName", value)
}

// GetVPortName gets the value of VPortName for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyVPortName() (value string, err error) {
	retValue, err := instance.GetProperty("VPortName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVPortState sets the value of VPortState for the instance
func (instance *MSFT_NetAdapterVPortSettingData) SetPropertyVPortState(value uint32) (err error) {
	return instance.SetProperty("VPortState", value)
}

// GetVPortState gets the value of VPortState for the instance
func (instance *MSFT_NetAdapterVPortSettingData) GetPropertyVPortState() (value uint32, err error) {
	retValue, err := instance.GetProperty("VPortState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
