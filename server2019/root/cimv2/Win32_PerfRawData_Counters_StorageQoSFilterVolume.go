// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_StorageQoSFilterVolume struct
type Win32_PerfRawData_Counters_StorageQoSFilterVolume struct {
	Win32_PerfRawData

	//
	AllocationQuantum uint64

	//
	AvgBandwidth uint64

	//
	AvgDeviceLatency uint64

	//
	AvgDeviceQueueLength uint64

	//
	AvgIOCost uint64

	//
	AvgNormalizedIOCost uint64

	//
	AvgSchedulerQueueLength uint64

	//
	CongestionThreshold uint64

	//
	DelayedCost uint64

	//
	EstimatedCapacity uint64

	//
	FlowSwitchCost uint64

	//
	IssuedCost uint64

	//
	LatencyTarget uint64

	//
	LowerThreshold uint64

	//
	NormalizedThroughput uint64

	//
	OverheadCost uint64

	//
	SectorCost uint64

	//
	SeekCost uint64
}

// SetAllocationQuantum sets the value of AllocationQuantum for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAllocationQuantum(value uint64) (err error) {
	return instance.SetProperty("AllocationQuantum", value)
}

// GetAllocationQuantum gets the value of AllocationQuantum for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAllocationQuantum() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationQuantum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgBandwidth sets the value of AvgBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgBandwidth(value uint64) (err error) {
	return instance.SetProperty("AvgBandwidth", value)
}

// GetAvgBandwidth gets the value of AvgBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDeviceLatency sets the value of AvgDeviceLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgDeviceLatency(value uint64) (err error) {
	return instance.SetProperty("AvgDeviceLatency", value)
}

// GetAvgDeviceLatency gets the value of AvgDeviceLatency for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgDeviceLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDeviceLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDeviceQueueLength sets the value of AvgDeviceQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgDeviceQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDeviceQueueLength", value)
}

// GetAvgDeviceQueueLength gets the value of AvgDeviceQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgDeviceQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDeviceQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgIOCost sets the value of AvgIOCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgIOCost(value uint64) (err error) {
	return instance.SetProperty("AvgIOCost", value)
}

// GetAvgIOCost gets the value of AvgIOCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgIOCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgIOCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgNormalizedIOCost sets the value of AvgNormalizedIOCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgNormalizedIOCost(value uint64) (err error) {
	return instance.SetProperty("AvgNormalizedIOCost", value)
}

// GetAvgNormalizedIOCost gets the value of AvgNormalizedIOCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgNormalizedIOCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgNormalizedIOCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgSchedulerQueueLength sets the value of AvgSchedulerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyAvgSchedulerQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgSchedulerQueueLength", value)
}

// GetAvgSchedulerQueueLength gets the value of AvgSchedulerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyAvgSchedulerQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgSchedulerQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCongestionThreshold sets the value of CongestionThreshold for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyCongestionThreshold(value uint64) (err error) {
	return instance.SetProperty("CongestionThreshold", value)
}

// GetCongestionThreshold gets the value of CongestionThreshold for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyCongestionThreshold() (value uint64, err error) {
	retValue, err := instance.GetProperty("CongestionThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDelayedCost sets the value of DelayedCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyDelayedCost(value uint64) (err error) {
	return instance.SetProperty("DelayedCost", value)
}

// GetDelayedCost gets the value of DelayedCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyDelayedCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("DelayedCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEstimatedCapacity sets the value of EstimatedCapacity for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyEstimatedCapacity(value uint64) (err error) {
	return instance.SetProperty("EstimatedCapacity", value)
}

// GetEstimatedCapacity gets the value of EstimatedCapacity for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyEstimatedCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("EstimatedCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlowSwitchCost sets the value of FlowSwitchCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyFlowSwitchCost(value uint64) (err error) {
	return instance.SetProperty("FlowSwitchCost", value)
}

// GetFlowSwitchCost gets the value of FlowSwitchCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyFlowSwitchCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlowSwitchCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIssuedCost sets the value of IssuedCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyIssuedCost(value uint64) (err error) {
	return instance.SetProperty("IssuedCost", value)
}

// GetIssuedCost gets the value of IssuedCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyIssuedCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("IssuedCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLatencyTarget sets the value of LatencyTarget for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyLatencyTarget(value uint64) (err error) {
	return instance.SetProperty("LatencyTarget", value)
}

// GetLatencyTarget gets the value of LatencyTarget for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyLatencyTarget() (value uint64, err error) {
	retValue, err := instance.GetProperty("LatencyTarget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowerThreshold sets the value of LowerThreshold for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyLowerThreshold(value uint64) (err error) {
	return instance.SetProperty("LowerThreshold", value)
}

// GetLowerThreshold gets the value of LowerThreshold for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyLowerThreshold() (value uint64, err error) {
	retValue, err := instance.GetProperty("LowerThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalizedThroughput sets the value of NormalizedThroughput for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyNormalizedThroughput(value uint64) (err error) {
	return instance.SetProperty("NormalizedThroughput", value)
}

// GetNormalizedThroughput gets the value of NormalizedThroughput for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyNormalizedThroughput() (value uint64, err error) {
	retValue, err := instance.GetProperty("NormalizedThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOverheadCost sets the value of OverheadCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertyOverheadCost(value uint64) (err error) {
	return instance.SetProperty("OverheadCost", value)
}

// GetOverheadCost gets the value of OverheadCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertyOverheadCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("OverheadCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSectorCost sets the value of SectorCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertySectorCost(value uint64) (err error) {
	return instance.SetProperty("SectorCost", value)
}

// GetSectorCost gets the value of SectorCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertySectorCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("SectorCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSeekCost sets the value of SeekCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) SetPropertySeekCost(value uint64) (err error) {
	return instance.SetProperty("SeekCost", value)
}

// GetSeekCost gets the value of SeekCost for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterVolume) GetPropertySeekCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("SeekCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
