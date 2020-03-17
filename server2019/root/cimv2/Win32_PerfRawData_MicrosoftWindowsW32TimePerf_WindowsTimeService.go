// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService struct
type Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService struct {
	Win32_PerfRawData

	//
	ClockFrequencyAdjustment uint32

	//
	ClockFrequencyAdjustmentPPB uint32

	//
	ComputedTimeOffset uint64

	//
	NTPClientTimeSourceCount uint32

	//
	NTPRoundtripDelay uint32

	//
	NTPServerIncomingRequests uint64

	//
	NTPServerOutgoingResponses uint64
}

// SetClockFrequencyAdjustment sets the value of ClockFrequencyAdjustment for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyClockFrequencyAdjustment(value uint32) (err error) {
	return instance.SetProperty("ClockFrequencyAdjustment", value)
}

// GetClockFrequencyAdjustment gets the value of ClockFrequencyAdjustment for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyClockFrequencyAdjustment() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClockFrequencyAdjustment")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClockFrequencyAdjustmentPPB sets the value of ClockFrequencyAdjustmentPPB for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyClockFrequencyAdjustmentPPB(value uint32) (err error) {
	return instance.SetProperty("ClockFrequencyAdjustmentPPB", value)
}

// GetClockFrequencyAdjustmentPPB gets the value of ClockFrequencyAdjustmentPPB for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyClockFrequencyAdjustmentPPB() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClockFrequencyAdjustmentPPB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputedTimeOffset sets the value of ComputedTimeOffset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyComputedTimeOffset(value uint64) (err error) {
	return instance.SetProperty("ComputedTimeOffset", value)
}

// GetComputedTimeOffset gets the value of ComputedTimeOffset for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyComputedTimeOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("ComputedTimeOffset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNTPClientTimeSourceCount sets the value of NTPClientTimeSourceCount for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyNTPClientTimeSourceCount(value uint32) (err error) {
	return instance.SetProperty("NTPClientTimeSourceCount", value)
}

// GetNTPClientTimeSourceCount gets the value of NTPClientTimeSourceCount for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyNTPClientTimeSourceCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NTPClientTimeSourceCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNTPRoundtripDelay sets the value of NTPRoundtripDelay for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyNTPRoundtripDelay(value uint32) (err error) {
	return instance.SetProperty("NTPRoundtripDelay", value)
}

// GetNTPRoundtripDelay gets the value of NTPRoundtripDelay for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyNTPRoundtripDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("NTPRoundtripDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNTPServerIncomingRequests sets the value of NTPServerIncomingRequests for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyNTPServerIncomingRequests(value uint64) (err error) {
	return instance.SetProperty("NTPServerIncomingRequests", value)
}

// GetNTPServerIncomingRequests gets the value of NTPServerIncomingRequests for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyNTPServerIncomingRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("NTPServerIncomingRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNTPServerOutgoingResponses sets the value of NTPServerOutgoingResponses for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) SetPropertyNTPServerOutgoingResponses(value uint64) (err error) {
	return instance.SetProperty("NTPServerOutgoingResponses", value)
}

// GetNTPServerOutgoingResponses gets the value of NTPServerOutgoingResponses for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsW32TimePerf_WindowsTimeService) GetPropertyNTPServerOutgoingResponses() (value uint64, err error) {
	retValue, err := instance.GetProperty("NTPServerOutgoingResponses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
