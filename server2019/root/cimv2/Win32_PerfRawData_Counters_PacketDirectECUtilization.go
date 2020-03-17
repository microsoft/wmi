// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_PacketDirectECUtilization struct
type Win32_PerfRawData_Counters_PacketDirectECUtilization struct {
	Win32_PerfRawData

	//
	BusyWaitIterationsPersec uint32

	//
	IterationsPersec uint32

	//
	PercentBusyWaitingTime uint64

	//
	PercentBusyWaitingTime_Base uint64

	//
	PercentBusyWaitIterations uint32

	//
	PercentBusyWaitIterations_Base uint32

	//
	PercentIdleTime uint64

	//
	PercentIdleTime_Base uint64

	//
	PercentProcessingTime uint64

	//
	PercentProcessingTime_Base uint64

	//
	ProcessorNumber uint32

	//
	RXQueueCount uint32

	//
	TotalBusyWaitIterations uint64

	//
	TotalIterations uint64

	//
	TXQueueCount uint32
}

// SetBusyWaitIterationsPersec sets the value of BusyWaitIterationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyBusyWaitIterationsPersec(value uint32) (err error) {
	return instance.SetProperty("BusyWaitIterationsPersec", value)
}

// GetBusyWaitIterationsPersec gets the value of BusyWaitIterationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyBusyWaitIterationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusyWaitIterationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIterationsPersec sets the value of IterationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyIterationsPersec(value uint32) (err error) {
	return instance.SetProperty("IterationsPersec", value)
}

// GetIterationsPersec gets the value of IterationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyIterationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IterationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentBusyWaitingTime sets the value of PercentBusyWaitingTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitingTime(value uint64) (err error) {
	return instance.SetProperty("PercentBusyWaitingTime", value)
}

// GetPercentBusyWaitingTime gets the value of PercentBusyWaitingTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitingTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentBusyWaitingTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentBusyWaitingTime_Base sets the value of PercentBusyWaitingTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitingTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentBusyWaitingTime_Base", value)
}

// GetPercentBusyWaitingTime_Base gets the value of PercentBusyWaitingTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitingTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentBusyWaitingTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentBusyWaitIterations sets the value of PercentBusyWaitIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitIterations(value uint32) (err error) {
	return instance.SetProperty("PercentBusyWaitIterations", value)
}

// GetPercentBusyWaitIterations gets the value of PercentBusyWaitIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitIterations() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentBusyWaitIterations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentBusyWaitIterations_Base sets the value of PercentBusyWaitIterations_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentBusyWaitIterations_Base(value uint32) (err error) {
	return instance.SetProperty("PercentBusyWaitIterations_Base", value)
}

// GetPercentBusyWaitIterations_Base gets the value of PercentBusyWaitIterations_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentBusyWaitIterations_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentBusyWaitIterations_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentIdleTime sets the value of PercentIdleTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentIdleTime(value uint64) (err error) {
	return instance.SetProperty("PercentIdleTime", value)
}

// GetPercentIdleTime gets the value of PercentIdleTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentIdleTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentIdleTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentIdleTime_Base sets the value of PercentIdleTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentIdleTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentIdleTime_Base", value)
}

// GetPercentIdleTime_Base gets the value of PercentIdleTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentIdleTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentIdleTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentProcessingTime sets the value of PercentProcessingTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentProcessingTime(value uint64) (err error) {
	return instance.SetProperty("PercentProcessingTime", value)
}

// GetPercentProcessingTime gets the value of PercentProcessingTime for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentProcessingTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentProcessingTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentProcessingTime_Base sets the value of PercentProcessingTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyPercentProcessingTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentProcessingTime_Base", value)
}

// GetPercentProcessingTime_Base gets the value of PercentProcessingTime_Base for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyPercentProcessingTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentProcessingTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyProcessorNumber(value uint32) (err error) {
	return instance.SetProperty("ProcessorNumber", value)
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyProcessorNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRXQueueCount sets the value of RXQueueCount for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyRXQueueCount(value uint32) (err error) {
	return instance.SetProperty("RXQueueCount", value)
}

// GetRXQueueCount gets the value of RXQueueCount for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyRXQueueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXQueueCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalBusyWaitIterations sets the value of TotalBusyWaitIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyTotalBusyWaitIterations(value uint64) (err error) {
	return instance.SetProperty("TotalBusyWaitIterations", value)
}

// GetTotalBusyWaitIterations gets the value of TotalBusyWaitIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyTotalBusyWaitIterations() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBusyWaitIterations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalIterations sets the value of TotalIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyTotalIterations(value uint64) (err error) {
	return instance.SetProperty("TotalIterations", value)
}

// GetTotalIterations gets the value of TotalIterations for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyTotalIterations() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalIterations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTXQueueCount sets the value of TXQueueCount for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) SetPropertyTXQueueCount(value uint32) (err error) {
	return instance.SetProperty("TXQueueCount", value)
}

// GetTXQueueCount gets the value of TXQueueCount for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectECUtilization) GetPropertyTXQueueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("TXQueueCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
