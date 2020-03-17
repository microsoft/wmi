// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor struct
type Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor struct {
	Win32_PerfRawData

	//
	C1TransitionsPersec uint64

	//
	C2TransitionsPersec uint64

	//
	C3TransitionsPersec uint64

	//
	ContextSwitchesPersec uint64

	//
	Frequency uint64

	//
	HardwareInterruptsPersec uint64

	//
	HWPRequestMSRContextSwitchesPersec uint64

	//
	HypervisorBranchPredictorFlushesPersec uint64

	//
	HypervisorImmediateL1DataCacheFlushesPersec uint64

	//
	HypervisorL1DataCacheFlushesPersec uint64

	//
	HypervisorMicroarchitecturalBufferFlushesPersec uint64

	//
	InterProcessorInterruptsPersec uint64

	//
	InterProcessorInterruptsSentPersec uint64

	//
	MonitorTransitionCost uint64

	//
	ParkingStatus uint64

	//
	PercentC1Time uint64

	//
	PercentC1Time_Base uint64

	//
	PercentC2Time uint64

	//
	PercentC2Time_Base uint64

	//
	PercentC3Time uint64

	//
	PercentC3Time_Base uint64

	//
	PercentGuestRunTime uint64

	//
	PercentGuestRunTime_Base uint64

	//
	PercentHypervisorRunTime uint64

	//
	PercentHypervisorRunTime_Base uint64

	//
	PercentIdleTime uint64

	//
	PercentIdleTime_Base uint64

	//
	PercentofMaxFrequency uint64

	//
	PercentTotalRunTime uint64

	//
	PercentTotalRunTime_Base uint64

	//
	PostedInterruptNotificationsPersec uint64

	//
	ProcessorStateFlags uint64

	//
	RootVpIndex uint64

	//
	SchedulerInterruptsPersec uint64

	//
	TimerInterruptsPersec uint64

	//
	TotalInterruptsPersec uint64
}

// SetC1TransitionsPersec sets the value of C1TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyC1TransitionsPersec(value uint64) (err error) {
	return instance.SetProperty("C1TransitionsPersec", value)
}

// GetC1TransitionsPersec gets the value of C1TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyC1TransitionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("C1TransitionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetC2TransitionsPersec sets the value of C2TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyC2TransitionsPersec(value uint64) (err error) {
	return instance.SetProperty("C2TransitionsPersec", value)
}

// GetC2TransitionsPersec gets the value of C2TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyC2TransitionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("C2TransitionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetC3TransitionsPersec sets the value of C3TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyC3TransitionsPersec(value uint64) (err error) {
	return instance.SetProperty("C3TransitionsPersec", value)
}

// GetC3TransitionsPersec gets the value of C3TransitionsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyC3TransitionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("C3TransitionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContextSwitchesPersec sets the value of ContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyContextSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("ContextSwitchesPersec", value)
}

// GetContextSwitchesPersec gets the value of ContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyContextSwitchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ContextSwitchesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrequency sets the value of Frequency for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyFrequency(value uint64) (err error) {
	return instance.SetProperty("Frequency", value)
}

// GetFrequency gets the value of Frequency for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyFrequency() (value uint64, err error) {
	retValue, err := instance.GetProperty("Frequency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareInterruptsPersec sets the value of HardwareInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHardwareInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("HardwareInterruptsPersec", value)
}

// GetHardwareInterruptsPersec gets the value of HardwareInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHardwareInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HardwareInterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHWPRequestMSRContextSwitchesPersec sets the value of HWPRequestMSRContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHWPRequestMSRContextSwitchesPersec(value uint64) (err error) {
	return instance.SetProperty("HWPRequestMSRContextSwitchesPersec", value)
}

// GetHWPRequestMSRContextSwitchesPersec gets the value of HWPRequestMSRContextSwitchesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHWPRequestMSRContextSwitchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HWPRequestMSRContextSwitchesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHypervisorBranchPredictorFlushesPersec sets the value of HypervisorBranchPredictorFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHypervisorBranchPredictorFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("HypervisorBranchPredictorFlushesPersec", value)
}

// GetHypervisorBranchPredictorFlushesPersec gets the value of HypervisorBranchPredictorFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHypervisorBranchPredictorFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorBranchPredictorFlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHypervisorImmediateL1DataCacheFlushesPersec sets the value of HypervisorImmediateL1DataCacheFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHypervisorImmediateL1DataCacheFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("HypervisorImmediateL1DataCacheFlushesPersec", value)
}

// GetHypervisorImmediateL1DataCacheFlushesPersec gets the value of HypervisorImmediateL1DataCacheFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHypervisorImmediateL1DataCacheFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorImmediateL1DataCacheFlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHypervisorL1DataCacheFlushesPersec sets the value of HypervisorL1DataCacheFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHypervisorL1DataCacheFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("HypervisorL1DataCacheFlushesPersec", value)
}

// GetHypervisorL1DataCacheFlushesPersec gets the value of HypervisorL1DataCacheFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHypervisorL1DataCacheFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorL1DataCacheFlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHypervisorMicroarchitecturalBufferFlushesPersec sets the value of HypervisorMicroarchitecturalBufferFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyHypervisorMicroarchitecturalBufferFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("HypervisorMicroarchitecturalBufferFlushesPersec", value)
}

// GetHypervisorMicroarchitecturalBufferFlushesPersec gets the value of HypervisorMicroarchitecturalBufferFlushesPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyHypervisorMicroarchitecturalBufferFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorMicroarchitecturalBufferFlushesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterProcessorInterruptsPersec sets the value of InterProcessorInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyInterProcessorInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("InterProcessorInterruptsPersec", value)
}

// GetInterProcessorInterruptsPersec gets the value of InterProcessorInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyInterProcessorInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterProcessorInterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterProcessorInterruptsSentPersec sets the value of InterProcessorInterruptsSentPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyInterProcessorInterruptsSentPersec(value uint64) (err error) {
	return instance.SetProperty("InterProcessorInterruptsSentPersec", value)
}

// GetInterProcessorInterruptsSentPersec gets the value of InterProcessorInterruptsSentPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyInterProcessorInterruptsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterProcessorInterruptsSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitorTransitionCost sets the value of MonitorTransitionCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyMonitorTransitionCost(value uint64) (err error) {
	return instance.SetProperty("MonitorTransitionCost", value)
}

// GetMonitorTransitionCost gets the value of MonitorTransitionCost for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyMonitorTransitionCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonitorTransitionCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParkingStatus sets the value of ParkingStatus for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyParkingStatus(value uint64) (err error) {
	return instance.SetProperty("ParkingStatus", value)
}

// GetParkingStatus gets the value of ParkingStatus for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyParkingStatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("ParkingStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC1Time sets the value of PercentC1Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC1Time(value uint64) (err error) {
	return instance.SetProperty("PercentC1Time", value)
}

// GetPercentC1Time gets the value of PercentC1Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC1Time() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC1Time")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC1Time_Base sets the value of PercentC1Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC1Time_Base(value uint64) (err error) {
	return instance.SetProperty("PercentC1Time_Base", value)
}

// GetPercentC1Time_Base gets the value of PercentC1Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC1Time_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC1Time_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC2Time sets the value of PercentC2Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC2Time(value uint64) (err error) {
	return instance.SetProperty("PercentC2Time", value)
}

// GetPercentC2Time gets the value of PercentC2Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC2Time() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC2Time")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC2Time_Base sets the value of PercentC2Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC2Time_Base(value uint64) (err error) {
	return instance.SetProperty("PercentC2Time_Base", value)
}

// GetPercentC2Time_Base gets the value of PercentC2Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC2Time_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC2Time_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC3Time sets the value of PercentC3Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC3Time(value uint64) (err error) {
	return instance.SetProperty("PercentC3Time", value)
}

// GetPercentC3Time gets the value of PercentC3Time for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC3Time() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC3Time")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentC3Time_Base sets the value of PercentC3Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentC3Time_Base(value uint64) (err error) {
	return instance.SetProperty("PercentC3Time_Base", value)
}

// GetPercentC3Time_Base gets the value of PercentC3Time_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentC3Time_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentC3Time_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentGuestRunTime sets the value of PercentGuestRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentGuestRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRunTime", value)
}

// GetPercentGuestRunTime gets the value of PercentGuestRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentGuestRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentGuestRunTime_Base sets the value of PercentGuestRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentGuestRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentGuestRunTime_Base", value)
}

// GetPercentGuestRunTime_Base gets the value of PercentGuestRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentGuestRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentGuestRunTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentHypervisorRunTime sets the value of PercentHypervisorRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentHypervisorRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentHypervisorRunTime", value)
}

// GetPercentHypervisorRunTime gets the value of PercentHypervisorRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentHypervisorRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentHypervisorRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentHypervisorRunTime_Base sets the value of PercentHypervisorRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentHypervisorRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentHypervisorRunTime_Base", value)
}

// GetPercentHypervisorRunTime_Base gets the value of PercentHypervisorRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentHypervisorRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentHypervisorRunTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentIdleTime sets the value of PercentIdleTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentIdleTime(value uint64) (err error) {
	return instance.SetProperty("PercentIdleTime", value)
}

// GetPercentIdleTime gets the value of PercentIdleTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentIdleTime() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentIdleTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentIdleTime_Base", value)
}

// GetPercentIdleTime_Base gets the value of PercentIdleTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentIdleTime_Base() (value uint64, err error) {
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

// SetPercentofMaxFrequency sets the value of PercentofMaxFrequency for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentofMaxFrequency(value uint64) (err error) {
	return instance.SetProperty("PercentofMaxFrequency", value)
}

// GetPercentofMaxFrequency gets the value of PercentofMaxFrequency for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentofMaxFrequency() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentofMaxFrequency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTotalRunTime sets the value of PercentTotalRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentTotalRunTime(value uint64) (err error) {
	return instance.SetProperty("PercentTotalRunTime", value)
}

// GetPercentTotalRunTime gets the value of PercentTotalRunTime for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentTotalRunTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTotalRunTime_Base sets the value of PercentTotalRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPercentTotalRunTime_Base(value uint64) (err error) {
	return instance.SetProperty("PercentTotalRunTime_Base", value)
}

// GetPercentTotalRunTime_Base gets the value of PercentTotalRunTime_Base for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPercentTotalRunTime_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTotalRunTime_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostedInterruptNotificationsPersec sets the value of PostedInterruptNotificationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyPostedInterruptNotificationsPersec(value uint64) (err error) {
	return instance.SetProperty("PostedInterruptNotificationsPersec", value)
}

// GetPostedInterruptNotificationsPersec gets the value of PostedInterruptNotificationsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyPostedInterruptNotificationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostedInterruptNotificationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorStateFlags sets the value of ProcessorStateFlags for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyProcessorStateFlags(value uint64) (err error) {
	return instance.SetProperty("ProcessorStateFlags", value)
}

// GetProcessorStateFlags gets the value of ProcessorStateFlags for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyProcessorStateFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessorStateFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRootVpIndex sets the value of RootVpIndex for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyRootVpIndex(value uint64) (err error) {
	return instance.SetProperty("RootVpIndex", value)
}

// GetRootVpIndex gets the value of RootVpIndex for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyRootVpIndex() (value uint64, err error) {
	retValue, err := instance.GetProperty("RootVpIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSchedulerInterruptsPersec sets the value of SchedulerInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertySchedulerInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("SchedulerInterruptsPersec", value)
}

// GetSchedulerInterruptsPersec gets the value of SchedulerInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertySchedulerInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SchedulerInterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimerInterruptsPersec sets the value of TimerInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyTimerInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("TimerInterruptsPersec", value)
}

// GetTimerInterruptsPersec gets the value of TimerInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyTimerInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimerInterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalInterruptsPersec sets the value of TotalInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) SetPropertyTotalInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("TotalInterruptsPersec", value)
}

// GetTotalInterruptsPersec gets the value of TotalInterruptsPersec for the instance
func (instance *Win32_PerfRawData_HvStats_HyperVHypervisorLogicalProcessor) GetPropertyTotalInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
