// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSClientSetting struct
type Win32_TSClientSetting struct {
	Win32_TerminalSetting

	//
	AdvancedRemoteAppGraphics uint32

	//
	AudioCaptureRedir uint32

	//
	AudioMapping uint32

	//
	AVC444ModePreferred uint32

	//
	ClipboardMapping uint32

	//
	ColorDepth uint32

	//
	ColorDepthPolicy uint32

	//
	COMPortMapping uint32

	//
	ConnectClientDrivesAtLogon uint32

	//
	ConnectionPolicy uint32

	//
	ConnectPrinterAtLogon uint32

	//
	DefaultToClientPrinter uint32

	//
	DriveMapping uint32

	//
	EncodeImageQuality uint32

	//
	HardwareGraphicsAdapter uint32

	//
	LPTPortMapping uint32

	//
	MaxMonitors uint32

	//
	MaxXResolution uint32

	//
	MaxYResolution uint32

	//
	PNPRedirection uint32

	//
	PolicyAdvancedRemoteAppGraphics uint32

	//
	PolicySourceAudioCaptureRedir uint32

	//
	PolicySourceAudioMapping uint32

	//
	PolicySourceAvc444ModePreferred uint32

	//
	PolicySourceClipboardMapping uint32

	//
	PolicySourceColorDepth uint32

	//
	PolicySourceColorDepthPolicy uint32

	//
	PolicySourceCOMPortMapping uint32

	//
	PolicySourceDefaultToClientPrinter uint32

	//
	PolicySourceDriveMapping uint32

	//
	PolicySourceEncodeImageQuality uint32

	//
	PolicySourceHardwareGraphicsAdapter uint32

	//
	PolicySourceLPTPortMapping uint32

	//
	PolicySourceMaxMonitors uint32

	//
	PolicySourceMaxResolution uint32

	//
	PolicySourcePNPRedirection uint32

	//
	PolicySourceRemoteSessionProfile uint32

	//
	PolicySourceSelectNetworkDetect uint32

	//
	PolicySourceSelectTransport uint32

	//
	PolicySourceVideoPlaybackRedir uint32

	//
	PolicySourceWindowsPrinterMapping uint32

	//
	RemoteSessionProfile uint32

	//
	SelectNetworkDetect uint32

	//
	SelectTransport uint32

	//
	VideoPlaybackRedir uint32

	//
	WindowsPrinterMapping uint32
}

// SetAdvancedRemoteAppGraphics sets the value of AdvancedRemoteAppGraphics for the instance
func (instance *Win32_TSClientSetting) SetPropertyAdvancedRemoteAppGraphics(value uint32) (err error) {
	return instance.SetProperty("AdvancedRemoteAppGraphics", value)
}

// GetAdvancedRemoteAppGraphics gets the value of AdvancedRemoteAppGraphics for the instance
func (instance *Win32_TSClientSetting) GetPropertyAdvancedRemoteAppGraphics() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdvancedRemoteAppGraphics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAudioCaptureRedir sets the value of AudioCaptureRedir for the instance
func (instance *Win32_TSClientSetting) SetPropertyAudioCaptureRedir(value uint32) (err error) {
	return instance.SetProperty("AudioCaptureRedir", value)
}

// GetAudioCaptureRedir gets the value of AudioCaptureRedir for the instance
func (instance *Win32_TSClientSetting) GetPropertyAudioCaptureRedir() (value uint32, err error) {
	retValue, err := instance.GetProperty("AudioCaptureRedir")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAudioMapping sets the value of AudioMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyAudioMapping(value uint32) (err error) {
	return instance.SetProperty("AudioMapping", value)
}

// GetAudioMapping gets the value of AudioMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyAudioMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("AudioMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAVC444ModePreferred sets the value of AVC444ModePreferred for the instance
func (instance *Win32_TSClientSetting) SetPropertyAVC444ModePreferred(value uint32) (err error) {
	return instance.SetProperty("AVC444ModePreferred", value)
}

// GetAVC444ModePreferred gets the value of AVC444ModePreferred for the instance
func (instance *Win32_TSClientSetting) GetPropertyAVC444ModePreferred() (value uint32, err error) {
	retValue, err := instance.GetProperty("AVC444ModePreferred")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClipboardMapping sets the value of ClipboardMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyClipboardMapping(value uint32) (err error) {
	return instance.SetProperty("ClipboardMapping", value)
}

// GetClipboardMapping gets the value of ClipboardMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyClipboardMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClipboardMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColorDepth sets the value of ColorDepth for the instance
func (instance *Win32_TSClientSetting) SetPropertyColorDepth(value uint32) (err error) {
	return instance.SetProperty("ColorDepth", value)
}

// GetColorDepth gets the value of ColorDepth for the instance
func (instance *Win32_TSClientSetting) GetPropertyColorDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("ColorDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColorDepthPolicy sets the value of ColorDepthPolicy for the instance
func (instance *Win32_TSClientSetting) SetPropertyColorDepthPolicy(value uint32) (err error) {
	return instance.SetProperty("ColorDepthPolicy", value)
}

// GetColorDepthPolicy gets the value of ColorDepthPolicy for the instance
func (instance *Win32_TSClientSetting) GetPropertyColorDepthPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ColorDepthPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCOMPortMapping sets the value of COMPortMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyCOMPortMapping(value uint32) (err error) {
	return instance.SetProperty("COMPortMapping", value)
}

// GetCOMPortMapping gets the value of COMPortMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyCOMPortMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("COMPortMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectClientDrivesAtLogon sets the value of ConnectClientDrivesAtLogon for the instance
func (instance *Win32_TSClientSetting) SetPropertyConnectClientDrivesAtLogon(value uint32) (err error) {
	return instance.SetProperty("ConnectClientDrivesAtLogon", value)
}

// GetConnectClientDrivesAtLogon gets the value of ConnectClientDrivesAtLogon for the instance
func (instance *Win32_TSClientSetting) GetPropertyConnectClientDrivesAtLogon() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectClientDrivesAtLogon")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionPolicy sets the value of ConnectionPolicy for the instance
func (instance *Win32_TSClientSetting) SetPropertyConnectionPolicy(value uint32) (err error) {
	return instance.SetProperty("ConnectionPolicy", value)
}

// GetConnectionPolicy gets the value of ConnectionPolicy for the instance
func (instance *Win32_TSClientSetting) GetPropertyConnectionPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectPrinterAtLogon sets the value of ConnectPrinterAtLogon for the instance
func (instance *Win32_TSClientSetting) SetPropertyConnectPrinterAtLogon(value uint32) (err error) {
	return instance.SetProperty("ConnectPrinterAtLogon", value)
}

// GetConnectPrinterAtLogon gets the value of ConnectPrinterAtLogon for the instance
func (instance *Win32_TSClientSetting) GetPropertyConnectPrinterAtLogon() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectPrinterAtLogon")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultToClientPrinter sets the value of DefaultToClientPrinter for the instance
func (instance *Win32_TSClientSetting) SetPropertyDefaultToClientPrinter(value uint32) (err error) {
	return instance.SetProperty("DefaultToClientPrinter", value)
}

// GetDefaultToClientPrinter gets the value of DefaultToClientPrinter for the instance
func (instance *Win32_TSClientSetting) GetPropertyDefaultToClientPrinter() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultToClientPrinter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriveMapping sets the value of DriveMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyDriveMapping(value uint32) (err error) {
	return instance.SetProperty("DriveMapping", value)
}

// GetDriveMapping gets the value of DriveMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyDriveMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncodeImageQuality sets the value of EncodeImageQuality for the instance
func (instance *Win32_TSClientSetting) SetPropertyEncodeImageQuality(value uint32) (err error) {
	return instance.SetProperty("EncodeImageQuality", value)
}

// GetEncodeImageQuality gets the value of EncodeImageQuality for the instance
func (instance *Win32_TSClientSetting) GetPropertyEncodeImageQuality() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncodeImageQuality")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareGraphicsAdapter sets the value of HardwareGraphicsAdapter for the instance
func (instance *Win32_TSClientSetting) SetPropertyHardwareGraphicsAdapter(value uint32) (err error) {
	return instance.SetProperty("HardwareGraphicsAdapter", value)
}

// GetHardwareGraphicsAdapter gets the value of HardwareGraphicsAdapter for the instance
func (instance *Win32_TSClientSetting) GetPropertyHardwareGraphicsAdapter() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareGraphicsAdapter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLPTPortMapping sets the value of LPTPortMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyLPTPortMapping(value uint32) (err error) {
	return instance.SetProperty("LPTPortMapping", value)
}

// GetLPTPortMapping gets the value of LPTPortMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyLPTPortMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("LPTPortMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxMonitors sets the value of MaxMonitors for the instance
func (instance *Win32_TSClientSetting) SetPropertyMaxMonitors(value uint32) (err error) {
	return instance.SetProperty("MaxMonitors", value)
}

// GetMaxMonitors gets the value of MaxMonitors for the instance
func (instance *Win32_TSClientSetting) GetPropertyMaxMonitors() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMonitors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxXResolution sets the value of MaxXResolution for the instance
func (instance *Win32_TSClientSetting) SetPropertyMaxXResolution(value uint32) (err error) {
	return instance.SetProperty("MaxXResolution", value)
}

// GetMaxXResolution gets the value of MaxXResolution for the instance
func (instance *Win32_TSClientSetting) GetPropertyMaxXResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxXResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxYResolution sets the value of MaxYResolution for the instance
func (instance *Win32_TSClientSetting) SetPropertyMaxYResolution(value uint32) (err error) {
	return instance.SetProperty("MaxYResolution", value)
}

// GetMaxYResolution gets the value of MaxYResolution for the instance
func (instance *Win32_TSClientSetting) GetPropertyMaxYResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxYResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPNPRedirection sets the value of PNPRedirection for the instance
func (instance *Win32_TSClientSetting) SetPropertyPNPRedirection(value uint32) (err error) {
	return instance.SetProperty("PNPRedirection", value)
}

// GetPNPRedirection gets the value of PNPRedirection for the instance
func (instance *Win32_TSClientSetting) GetPropertyPNPRedirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("PNPRedirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyAdvancedRemoteAppGraphics sets the value of PolicyAdvancedRemoteAppGraphics for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicyAdvancedRemoteAppGraphics(value uint32) (err error) {
	return instance.SetProperty("PolicyAdvancedRemoteAppGraphics", value)
}

// GetPolicyAdvancedRemoteAppGraphics gets the value of PolicyAdvancedRemoteAppGraphics for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicyAdvancedRemoteAppGraphics() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicyAdvancedRemoteAppGraphics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceAudioCaptureRedir sets the value of PolicySourceAudioCaptureRedir for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceAudioCaptureRedir(value uint32) (err error) {
	return instance.SetProperty("PolicySourceAudioCaptureRedir", value)
}

// GetPolicySourceAudioCaptureRedir gets the value of PolicySourceAudioCaptureRedir for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceAudioCaptureRedir() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceAudioCaptureRedir")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceAudioMapping sets the value of PolicySourceAudioMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceAudioMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceAudioMapping", value)
}

// GetPolicySourceAudioMapping gets the value of PolicySourceAudioMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceAudioMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceAudioMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceAvc444ModePreferred sets the value of PolicySourceAvc444ModePreferred for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceAvc444ModePreferred(value uint32) (err error) {
	return instance.SetProperty("PolicySourceAvc444ModePreferred", value)
}

// GetPolicySourceAvc444ModePreferred gets the value of PolicySourceAvc444ModePreferred for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceAvc444ModePreferred() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceAvc444ModePreferred")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceClipboardMapping sets the value of PolicySourceClipboardMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceClipboardMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceClipboardMapping", value)
}

// GetPolicySourceClipboardMapping gets the value of PolicySourceClipboardMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceClipboardMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceClipboardMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceColorDepth sets the value of PolicySourceColorDepth for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceColorDepth(value uint32) (err error) {
	return instance.SetProperty("PolicySourceColorDepth", value)
}

// GetPolicySourceColorDepth gets the value of PolicySourceColorDepth for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceColorDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceColorDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceColorDepthPolicy sets the value of PolicySourceColorDepthPolicy for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceColorDepthPolicy(value uint32) (err error) {
	return instance.SetProperty("PolicySourceColorDepthPolicy", value)
}

// GetPolicySourceColorDepthPolicy gets the value of PolicySourceColorDepthPolicy for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceColorDepthPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceColorDepthPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceCOMPortMapping sets the value of PolicySourceCOMPortMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceCOMPortMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceCOMPortMapping", value)
}

// GetPolicySourceCOMPortMapping gets the value of PolicySourceCOMPortMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceCOMPortMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceCOMPortMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceDefaultToClientPrinter sets the value of PolicySourceDefaultToClientPrinter for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceDefaultToClientPrinter(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDefaultToClientPrinter", value)
}

// GetPolicySourceDefaultToClientPrinter gets the value of PolicySourceDefaultToClientPrinter for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceDefaultToClientPrinter() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDefaultToClientPrinter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceDriveMapping sets the value of PolicySourceDriveMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceDriveMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDriveMapping", value)
}

// GetPolicySourceDriveMapping gets the value of PolicySourceDriveMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceDriveMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDriveMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceEncodeImageQuality sets the value of PolicySourceEncodeImageQuality for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceEncodeImageQuality(value uint32) (err error) {
	return instance.SetProperty("PolicySourceEncodeImageQuality", value)
}

// GetPolicySourceEncodeImageQuality gets the value of PolicySourceEncodeImageQuality for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceEncodeImageQuality() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceEncodeImageQuality")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceHardwareGraphicsAdapter sets the value of PolicySourceHardwareGraphicsAdapter for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceHardwareGraphicsAdapter(value uint32) (err error) {
	return instance.SetProperty("PolicySourceHardwareGraphicsAdapter", value)
}

// GetPolicySourceHardwareGraphicsAdapter gets the value of PolicySourceHardwareGraphicsAdapter for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceHardwareGraphicsAdapter() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceHardwareGraphicsAdapter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceLPTPortMapping sets the value of PolicySourceLPTPortMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceLPTPortMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceLPTPortMapping", value)
}

// GetPolicySourceLPTPortMapping gets the value of PolicySourceLPTPortMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceLPTPortMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceLPTPortMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceMaxMonitors sets the value of PolicySourceMaxMonitors for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceMaxMonitors(value uint32) (err error) {
	return instance.SetProperty("PolicySourceMaxMonitors", value)
}

// GetPolicySourceMaxMonitors gets the value of PolicySourceMaxMonitors for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceMaxMonitors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceMaxMonitors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceMaxResolution sets the value of PolicySourceMaxResolution for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceMaxResolution(value uint32) (err error) {
	return instance.SetProperty("PolicySourceMaxResolution", value)
}

// GetPolicySourceMaxResolution gets the value of PolicySourceMaxResolution for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceMaxResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceMaxResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourcePNPRedirection sets the value of PolicySourcePNPRedirection for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourcePNPRedirection(value uint32) (err error) {
	return instance.SetProperty("PolicySourcePNPRedirection", value)
}

// GetPolicySourcePNPRedirection gets the value of PolicySourcePNPRedirection for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourcePNPRedirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourcePNPRedirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceRemoteSessionProfile sets the value of PolicySourceRemoteSessionProfile for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceRemoteSessionProfile(value uint32) (err error) {
	return instance.SetProperty("PolicySourceRemoteSessionProfile", value)
}

// GetPolicySourceRemoteSessionProfile gets the value of PolicySourceRemoteSessionProfile for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceRemoteSessionProfile() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceRemoteSessionProfile")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceSelectNetworkDetect sets the value of PolicySourceSelectNetworkDetect for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceSelectNetworkDetect(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSelectNetworkDetect", value)
}

// GetPolicySourceSelectNetworkDetect gets the value of PolicySourceSelectNetworkDetect for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceSelectNetworkDetect() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSelectNetworkDetect")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceSelectTransport sets the value of PolicySourceSelectTransport for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceSelectTransport(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSelectTransport", value)
}

// GetPolicySourceSelectTransport gets the value of PolicySourceSelectTransport for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceSelectTransport() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSelectTransport")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceVideoPlaybackRedir sets the value of PolicySourceVideoPlaybackRedir for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceVideoPlaybackRedir(value uint32) (err error) {
	return instance.SetProperty("PolicySourceVideoPlaybackRedir", value)
}

// GetPolicySourceVideoPlaybackRedir gets the value of PolicySourceVideoPlaybackRedir for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceVideoPlaybackRedir() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceVideoPlaybackRedir")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceWindowsPrinterMapping sets the value of PolicySourceWindowsPrinterMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyPolicySourceWindowsPrinterMapping(value uint32) (err error) {
	return instance.SetProperty("PolicySourceWindowsPrinterMapping", value)
}

// GetPolicySourceWindowsPrinterMapping gets the value of PolicySourceWindowsPrinterMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyPolicySourceWindowsPrinterMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceWindowsPrinterMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteSessionProfile sets the value of RemoteSessionProfile for the instance
func (instance *Win32_TSClientSetting) SetPropertyRemoteSessionProfile(value uint32) (err error) {
	return instance.SetProperty("RemoteSessionProfile", value)
}

// GetRemoteSessionProfile gets the value of RemoteSessionProfile for the instance
func (instance *Win32_TSClientSetting) GetPropertyRemoteSessionProfile() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemoteSessionProfile")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectNetworkDetect sets the value of SelectNetworkDetect for the instance
func (instance *Win32_TSClientSetting) SetPropertySelectNetworkDetect(value uint32) (err error) {
	return instance.SetProperty("SelectNetworkDetect", value)
}

// GetSelectNetworkDetect gets the value of SelectNetworkDetect for the instance
func (instance *Win32_TSClientSetting) GetPropertySelectNetworkDetect() (value uint32, err error) {
	retValue, err := instance.GetProperty("SelectNetworkDetect")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectTransport sets the value of SelectTransport for the instance
func (instance *Win32_TSClientSetting) SetPropertySelectTransport(value uint32) (err error) {
	return instance.SetProperty("SelectTransport", value)
}

// GetSelectTransport gets the value of SelectTransport for the instance
func (instance *Win32_TSClientSetting) GetPropertySelectTransport() (value uint32, err error) {
	retValue, err := instance.GetProperty("SelectTransport")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVideoPlaybackRedir sets the value of VideoPlaybackRedir for the instance
func (instance *Win32_TSClientSetting) SetPropertyVideoPlaybackRedir(value uint32) (err error) {
	return instance.SetProperty("VideoPlaybackRedir", value)
}

// GetVideoPlaybackRedir gets the value of VideoPlaybackRedir for the instance
func (instance *Win32_TSClientSetting) GetPropertyVideoPlaybackRedir() (value uint32, err error) {
	retValue, err := instance.GetProperty("VideoPlaybackRedir")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWindowsPrinterMapping sets the value of WindowsPrinterMapping for the instance
func (instance *Win32_TSClientSetting) SetPropertyWindowsPrinterMapping(value uint32) (err error) {
	return instance.SetProperty("WindowsPrinterMapping", value)
}

// GetWindowsPrinterMapping gets the value of WindowsPrinterMapping for the instance
func (instance *Win32_TSClientSetting) GetPropertyWindowsPrinterMapping() (value uint32, err error) {
	retValue, err := instance.GetProperty("WindowsPrinterMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ColorDepthPolicy" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetColorDepthPolicy( /* IN */ ColorDepthPolicy uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetColorDepthPolicy", ColorDepthPolicy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ColorDepth" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetColorDepth( /* IN */ ColorDepth uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetColorDepth", ColorDepth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MaxMonitors" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetMaxMonitors( /* IN */ MaxMonitors uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetMaxMonitors", MaxMonitors)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MaxXResolution" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetMaxXResolution( /* IN */ MaxXResolution uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetMaxXResolution", MaxXResolution)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MaxYResolution" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetMaxYResolution( /* IN */ MaxYResolution uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetMaxYResolution", MaxYResolution)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ConnectClientDrivesAtLogon" type="uint32 "></param>
// <param name="ConnectPrinterAtLogon" type="uint32 "></param>
// <param name="DefaultToClientPrinter" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) ConnectionSettings( /* IN */ ConnectClientDrivesAtLogon uint32,
	/* IN */ ConnectPrinterAtLogon uint32,
	/* IN */ DefaultToClientPrinter uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ConnectionSettings", ConnectClientDrivesAtLogon, ConnectPrinterAtLogon, DefaultToClientPrinter)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="PropertyName" type="string "></param>
// <param name="Value" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSClientSetting) SetClientProperty( /* IN */ PropertyName string,
	/* IN */ Value bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetClientProperty", PropertyName, Value)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
