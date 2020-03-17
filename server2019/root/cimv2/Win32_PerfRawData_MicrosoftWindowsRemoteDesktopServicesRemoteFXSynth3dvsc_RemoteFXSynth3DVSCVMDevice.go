// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice struct
type Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice struct {
	Win32_PerfRawData

	//
	NumberofconnectedVMTchannels uint32

	//
	NumberofcreatedVMTchannels uint32

	//
	NumberofdisconnectedVMTchannels uint32

	//
	NumberofRDVGMrestartednotifications uint32

	//
	NumberofwaitingVMTchannels uint32

	//
	TotalnumberofcreatedVMTchannels uint32
}

// SetNumberofconnectedVMTchannels sets the value of NumberofconnectedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyNumberofconnectedVMTchannels(value uint32) (err error) {
	return instance.SetProperty("NumberofconnectedVMTchannels", value)
}

// GetNumberofconnectedVMTchannels gets the value of NumberofconnectedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyNumberofconnectedVMTchannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofconnectedVMTchannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofcreatedVMTchannels sets the value of NumberofcreatedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyNumberofcreatedVMTchannels(value uint32) (err error) {
	return instance.SetProperty("NumberofcreatedVMTchannels", value)
}

// GetNumberofcreatedVMTchannels gets the value of NumberofcreatedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyNumberofcreatedVMTchannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofcreatedVMTchannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofdisconnectedVMTchannels sets the value of NumberofdisconnectedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyNumberofdisconnectedVMTchannels(value uint32) (err error) {
	return instance.SetProperty("NumberofdisconnectedVMTchannels", value)
}

// GetNumberofdisconnectedVMTchannels gets the value of NumberofdisconnectedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyNumberofdisconnectedVMTchannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofdisconnectedVMTchannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofRDVGMrestartednotifications sets the value of NumberofRDVGMrestartednotifications for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyNumberofRDVGMrestartednotifications(value uint32) (err error) {
	return instance.SetProperty("NumberofRDVGMrestartednotifications", value)
}

// GetNumberofRDVGMrestartednotifications gets the value of NumberofRDVGMrestartednotifications for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyNumberofRDVGMrestartednotifications() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofRDVGMrestartednotifications")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofwaitingVMTchannels sets the value of NumberofwaitingVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyNumberofwaitingVMTchannels(value uint32) (err error) {
	return instance.SetProperty("NumberofwaitingVMTchannels", value)
}

// GetNumberofwaitingVMTchannels gets the value of NumberofwaitingVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyNumberofwaitingVMTchannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofwaitingVMTchannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalnumberofcreatedVMTchannels sets the value of TotalnumberofcreatedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) SetPropertyTotalnumberofcreatedVMTchannels(value uint32) (err error) {
	return instance.SetProperty("TotalnumberofcreatedVMTchannels", value)
}

// GetTotalnumberofcreatedVMTchannels gets the value of TotalnumberofcreatedVMTchannels for the instance
func (instance *Win32_PerfRawData_MicrosoftWindowsRemoteDesktopServicesRemoteFXSynth3dvsc_RemoteFXSynth3DVSCVMDevice) GetPropertyTotalnumberofcreatedVMTchannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalnumberofcreatedVMTchannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
