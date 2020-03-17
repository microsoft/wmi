// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_tapisrv_Telephony struct
type Win32_PerfFormattedData_tapisrv_Telephony struct {
	Win32_PerfFormattedData

	//
	ActiveLines uint32

	//
	ActiveTelephones uint32

	//
	ClientApps uint32

	//
	CurrentIncomingCalls uint32

	//
	CurrentOutgoingCalls uint32

	//
	IncomingCallsPersec uint32

	//
	Lines uint32

	//
	OutgoingCallsPersec uint32

	//
	TelephoneDevices uint32
}

// SetActiveLines sets the value of ActiveLines for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyActiveLines(value uint32) (err error) {
	return instance.SetProperty("ActiveLines", value)
}

// GetActiveLines gets the value of ActiveLines for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyActiveLines() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveLines")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveTelephones sets the value of ActiveTelephones for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyActiveTelephones(value uint32) (err error) {
	return instance.SetProperty("ActiveTelephones", value)
}

// GetActiveTelephones gets the value of ActiveTelephones for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyActiveTelephones() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveTelephones")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientApps sets the value of ClientApps for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyClientApps(value uint32) (err error) {
	return instance.SetProperty("ClientApps", value)
}

// GetClientApps gets the value of ClientApps for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyClientApps() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientApps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentIncomingCalls sets the value of CurrentIncomingCalls for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyCurrentIncomingCalls(value uint32) (err error) {
	return instance.SetProperty("CurrentIncomingCalls", value)
}

// GetCurrentIncomingCalls gets the value of CurrentIncomingCalls for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyCurrentIncomingCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentIncomingCalls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentOutgoingCalls sets the value of CurrentOutgoingCalls for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyCurrentOutgoingCalls(value uint32) (err error) {
	return instance.SetProperty("CurrentOutgoingCalls", value)
}

// GetCurrentOutgoingCalls gets the value of CurrentOutgoingCalls for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyCurrentOutgoingCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentOutgoingCalls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncomingCallsPersec sets the value of IncomingCallsPersec for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyIncomingCallsPersec(value uint32) (err error) {
	return instance.SetProperty("IncomingCallsPersec", value)
}

// GetIncomingCallsPersec gets the value of IncomingCallsPersec for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyIncomingCallsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IncomingCallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLines sets the value of Lines for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyLines(value uint32) (err error) {
	return instance.SetProperty("Lines", value)
}

// GetLines gets the value of Lines for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyLines() (value uint32, err error) {
	retValue, err := instance.GetProperty("Lines")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutgoingCallsPersec sets the value of OutgoingCallsPersec for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyOutgoingCallsPersec(value uint32) (err error) {
	return instance.SetProperty("OutgoingCallsPersec", value)
}

// GetOutgoingCallsPersec gets the value of OutgoingCallsPersec for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyOutgoingCallsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutgoingCallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTelephoneDevices sets the value of TelephoneDevices for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) SetPropertyTelephoneDevices(value uint32) (err error) {
	return instance.SetProperty("TelephoneDevices", value)
}

// GetTelephoneDevices gets the value of TelephoneDevices for the instance
func (instance *Win32_PerfFormattedData_tapisrv_Telephony) GetPropertyTelephoneDevices() (value uint32, err error) {
	retValue, err := instance.GetProperty("TelephoneDevices")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
