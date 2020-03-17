// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_XHCIInterrupter struct
type Win32_PerfRawData_Counters_XHCIInterrupter struct {
	Win32_PerfRawData

	//
	DpcRequeueCount uint32

	//
	DPCsPersec uint32

	//
	EventRingFullCount uint32

	//
	EventsprocessedDPC uint64

	//
	EventsprocessedDPC_Base uint32

	//
	InterruptsPersec uint32
}

// SetDpcRequeueCount sets the value of DpcRequeueCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyDpcRequeueCount(value uint32) (err error) {
	return instance.SetProperty("DpcRequeueCount", value)
}

// GetDpcRequeueCount gets the value of DpcRequeueCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyDpcRequeueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("DpcRequeueCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDPCsPersec sets the value of DPCsPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyDPCsPersec(value uint32) (err error) {
	return instance.SetProperty("DPCsPersec", value)
}

// GetDPCsPersec gets the value of DPCsPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyDPCsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DPCsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventRingFullCount sets the value of EventRingFullCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyEventRingFullCount(value uint32) (err error) {
	return instance.SetProperty("EventRingFullCount", value)
}

// GetEventRingFullCount gets the value of EventRingFullCount for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyEventRingFullCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventRingFullCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventsprocessedDPC sets the value of EventsprocessedDPC for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyEventsprocessedDPC(value uint64) (err error) {
	return instance.SetProperty("EventsprocessedDPC", value)
}

// GetEventsprocessedDPC gets the value of EventsprocessedDPC for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyEventsprocessedDPC() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventsprocessedDPC")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventsprocessedDPC_Base sets the value of EventsprocessedDPC_Base for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyEventsprocessedDPC_Base(value uint32) (err error) {
	return instance.SetProperty("EventsprocessedDPC_Base", value)
}

// GetEventsprocessedDPC_Base gets the value of EventsprocessedDPC_Base for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyEventsprocessedDPC_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventsprocessedDPC_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterruptsPersec sets the value of InterruptsPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) SetPropertyInterruptsPersec(value uint32) (err error) {
	return instance.SetProperty("InterruptsPersec", value)
}

// GetInterruptsPersec gets the value of InterruptsPersec for the instance
func (instance *Win32_PerfRawData_Counters_XHCIInterrupter) GetPropertyInterruptsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterruptsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
