// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic struct
type Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic struct {
	Win32_PerfFormattedData

	//
	AverageInboundBytesAllowedThroughtheQueue uint64

	//
	AverageInboundBytesDropped uint64

	//
	AverageInboundBytesEnteringtheQueue uint64

	//
	AverageInboundBytesQueuedduetoBacklog uint64

	//
	AverageInboundBytesQueuedduetoInsufficientTokens uint64

	//
	AverageInboundBytesResumed uint64

	//
	AverageInboundPacketsAllowedThroughtheQueue uint64

	//
	AverageInboundPacketsDropped uint64

	//
	AverageInboundPacketsEnteringtheQueue uint64

	//
	AverageInboundPacketsQueuedduetoBacklog uint64

	//
	AverageInboundPacketsQueuedduetoInsufficientTokens uint64

	//
	AverageInboundPacketsResumed uint64
}

// SetAverageInboundBytesAllowedThroughtheQueue sets the value of AverageInboundBytesAllowedThroughtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesAllowedThroughtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesAllowedThroughtheQueue", value)
}

// GetAverageInboundBytesAllowedThroughtheQueue gets the value of AverageInboundBytesAllowedThroughtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesAllowedThroughtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesAllowedThroughtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundBytesDropped sets the value of AverageInboundBytesDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesDropped(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesDropped", value)
}

// GetAverageInboundBytesDropped gets the value of AverageInboundBytesDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundBytesEnteringtheQueue sets the value of AverageInboundBytesEnteringtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesEnteringtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesEnteringtheQueue", value)
}

// GetAverageInboundBytesEnteringtheQueue gets the value of AverageInboundBytesEnteringtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesEnteringtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesEnteringtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundBytesQueuedduetoBacklog sets the value of AverageInboundBytesQueuedduetoBacklog for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesQueuedduetoBacklog(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesQueuedduetoBacklog", value)
}

// GetAverageInboundBytesQueuedduetoBacklog gets the value of AverageInboundBytesQueuedduetoBacklog for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesQueuedduetoBacklog() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesQueuedduetoBacklog")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundBytesQueuedduetoInsufficientTokens sets the value of AverageInboundBytesQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesQueuedduetoInsufficientTokens(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesQueuedduetoInsufficientTokens", value)
}

// GetAverageInboundBytesQueuedduetoInsufficientTokens gets the value of AverageInboundBytesQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesQueuedduetoInsufficientTokens() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesQueuedduetoInsufficientTokens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundBytesResumed sets the value of AverageInboundBytesResumed for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundBytesResumed(value uint64) (err error) {
	return instance.SetProperty("AverageInboundBytesResumed", value)
}

// GetAverageInboundBytesResumed gets the value of AverageInboundBytesResumed for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundBytesResumed() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundBytesResumed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsAllowedThroughtheQueue sets the value of AverageInboundPacketsAllowedThroughtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsAllowedThroughtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsAllowedThroughtheQueue", value)
}

// GetAverageInboundPacketsAllowedThroughtheQueue gets the value of AverageInboundPacketsAllowedThroughtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsAllowedThroughtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsAllowedThroughtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsDropped sets the value of AverageInboundPacketsDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsDropped(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsDropped", value)
}

// GetAverageInboundPacketsDropped gets the value of AverageInboundPacketsDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsEnteringtheQueue sets the value of AverageInboundPacketsEnteringtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsEnteringtheQueue(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsEnteringtheQueue", value)
}

// GetAverageInboundPacketsEnteringtheQueue gets the value of AverageInboundPacketsEnteringtheQueue for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsEnteringtheQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsEnteringtheQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsQueuedduetoBacklog sets the value of AverageInboundPacketsQueuedduetoBacklog for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsQueuedduetoBacklog(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsQueuedduetoBacklog", value)
}

// GetAverageInboundPacketsQueuedduetoBacklog gets the value of AverageInboundPacketsQueuedduetoBacklog for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsQueuedduetoBacklog() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsQueuedduetoBacklog")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsQueuedduetoInsufficientTokens sets the value of AverageInboundPacketsQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsQueuedduetoInsufficientTokens(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsQueuedduetoInsufficientTokens", value)
}

// GetAverageInboundPacketsQueuedduetoInsufficientTokens gets the value of AverageInboundPacketsQueuedduetoInsufficientTokens for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsQueuedduetoInsufficientTokens() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsQueuedduetoInsufficientTokens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageInboundPacketsResumed sets the value of AverageInboundPacketsResumed for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) SetPropertyAverageInboundPacketsResumed(value uint64) (err error) {
	return instance.SetProperty("AverageInboundPacketsResumed", value)
}

// GetAverageInboundPacketsResumed gets the value of AverageInboundPacketsResumed for the instance
func (instance *Win32_PerfFormattedData_Counters_VFPQoSQueueAverageInboundNetworkTraffic) GetPropertyAverageInboundPacketsResumed() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageInboundPacketsResumed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
