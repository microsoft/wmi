// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterVmqQueueSettingData struct
type MSFT_NetAdapterVmqQueueSettingData struct {
	MSFT_NetAdapterSettingData

	//
	FilterList []MSFT_NetAdapter_VmqFilter

	//
	NumFilters uint32

	//
	ProcessorAffinityMask uint64

	//
	ProcessorGroup uint16

	//
	QueueID uint32

	//
	QueueName string

	//
	State uint32

	//
	VmFriendlyName string

	//
	VmID string
}

// SetFilterList sets the value of FilterList for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyFilterList(value []MSFT_NetAdapter_VmqFilter) (err error) {
	return instance.SetProperty("FilterList", value)
}

// GetFilterList gets the value of FilterList for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyFilterList() (value []MSFT_NetAdapter_VmqFilter, err error) {
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

// SetNumFilters sets the value of NumFilters for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyNumFilters(value uint32) (err error) {
	return instance.SetProperty("NumFilters", value)
}

// GetNumFilters gets the value of NumFilters for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyNumFilters() (value uint32, err error) {
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

// SetProcessorAffinityMask sets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyProcessorAffinityMask(value uint64) (err error) {
	return instance.SetProperty("ProcessorAffinityMask", value)
}

// GetProcessorAffinityMask gets the value of ProcessorAffinityMask for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyProcessorAffinityMask() (value uint64, err error) {
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
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", value)
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyProcessorGroup() (value uint16, err error) {
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

// SetQueueID sets the value of QueueID for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyQueueID(value uint32) (err error) {
	return instance.SetProperty("QueueID", value)
}

// GetQueueID gets the value of QueueID for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyQueueID() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueueName sets the value of QueueName for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyQueueName(value string) (err error) {
	return instance.SetProperty("QueueName", value)
}

// GetQueueName gets the value of QueueName for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyQueueName() (value string, err error) {
	retValue, err := instance.GetProperty("QueueName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmFriendlyName sets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyVmFriendlyName(value string) (err error) {
	return instance.SetProperty("VmFriendlyName", value)
}

// GetVmFriendlyName gets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyVmFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("VmFriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmID sets the value of VmID for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) SetPropertyVmID(value string) (err error) {
	return instance.SetProperty("VmID", value)
}

// GetVmID gets the value of VmID for the instance
func (instance *MSFT_NetAdapterVmqQueueSettingData) GetPropertyVmID() (value string, err error) {
	retValue, err := instance.GetProperty("VmID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
