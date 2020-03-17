// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_DNS64Global struct
type Win32_PerfRawData_Counters_DNS64Global struct {
	Win32_PerfRawData

	//
	AAAAqueriesFailed uint64

	//
	AAAAqueriesSuccessful uint64

	//
	AAAASynthesizedrecords uint64

	//
	IP6ARPAqueriesMatched uint64

	//
	OtherqueriesFailed uint64

	//
	OtherqueriesSuccessful uint64
}

// SetAAAAqueriesFailed sets the value of AAAAqueriesFailed for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyAAAAqueriesFailed(value uint64) (err error) {
	return instance.SetProperty("AAAAqueriesFailed", value)
}

// GetAAAAqueriesFailed gets the value of AAAAqueriesFailed for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyAAAAqueriesFailed() (value uint64, err error) {
	retValue, err := instance.GetProperty("AAAAqueriesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAAAAqueriesSuccessful sets the value of AAAAqueriesSuccessful for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyAAAAqueriesSuccessful(value uint64) (err error) {
	return instance.SetProperty("AAAAqueriesSuccessful", value)
}

// GetAAAAqueriesSuccessful gets the value of AAAAqueriesSuccessful for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyAAAAqueriesSuccessful() (value uint64, err error) {
	retValue, err := instance.GetProperty("AAAAqueriesSuccessful")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAAAASynthesizedrecords sets the value of AAAASynthesizedrecords for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyAAAASynthesizedrecords(value uint64) (err error) {
	return instance.SetProperty("AAAASynthesizedrecords", value)
}

// GetAAAASynthesizedrecords gets the value of AAAASynthesizedrecords for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyAAAASynthesizedrecords() (value uint64, err error) {
	retValue, err := instance.GetProperty("AAAASynthesizedrecords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIP6ARPAqueriesMatched sets the value of IP6ARPAqueriesMatched for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyIP6ARPAqueriesMatched(value uint64) (err error) {
	return instance.SetProperty("IP6ARPAqueriesMatched", value)
}

// GetIP6ARPAqueriesMatched gets the value of IP6ARPAqueriesMatched for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyIP6ARPAqueriesMatched() (value uint64, err error) {
	retValue, err := instance.GetProperty("IP6ARPAqueriesMatched")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherqueriesFailed sets the value of OtherqueriesFailed for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyOtherqueriesFailed(value uint64) (err error) {
	return instance.SetProperty("OtherqueriesFailed", value)
}

// GetOtherqueriesFailed gets the value of OtherqueriesFailed for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyOtherqueriesFailed() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherqueriesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherqueriesSuccessful sets the value of OtherqueriesSuccessful for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) SetPropertyOtherqueriesSuccessful(value uint64) (err error) {
	return instance.SetProperty("OtherqueriesSuccessful", value)
}

// GetOtherqueriesSuccessful gets the value of OtherqueriesSuccessful for the instance
func (instance *Win32_PerfRawData_Counters_DNS64Global) GetPropertyOtherqueriesSuccessful() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherqueriesSuccessful")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
