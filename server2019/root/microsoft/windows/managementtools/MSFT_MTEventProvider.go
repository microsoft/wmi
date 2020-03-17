// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTEventProvider struct
type MSFT_MTEventProvider struct {
	CIM_ManagedElement

	//
	DisplayName string

	//
	DisplayPath string

	//
	ExportedChannelsCount uint32

	//
	Name string
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_MTEventProvider) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_MTEventProvider) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayPath sets the value of DisplayPath for the instance
func (instance *MSFT_MTEventProvider) SetPropertyDisplayPath(value string) (err error) {
	return instance.SetProperty("DisplayPath", value)
}

// GetDisplayPath gets the value of DisplayPath for the instance
func (instance *MSFT_MTEventProvider) GetPropertyDisplayPath() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedChannelsCount sets the value of ExportedChannelsCount for the instance
func (instance *MSFT_MTEventProvider) SetPropertyExportedChannelsCount(value uint32) (err error) {
	return instance.SetProperty("ExportedChannelsCount", value)
}

// GetExportedChannelsCount gets the value of ExportedChannelsCount for the instance
func (instance *MSFT_MTEventProvider) GetPropertyExportedChannelsCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExportedChannelsCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTEventProvider) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTEventProvider) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Result" type="MSFT_MTEventChannel []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventProvider) GetChannels( /* OUT */ Result []MSFT_MTEventChannel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetChannels")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EventProviders" type="MSFT_MTEventProvider []"></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="WindowsEventChannels" type="MSFT_MTEventChannel []"></param>
func (instance *MSFT_MTEventProvider) GetProvidersAndWindowsEventChannels( /* OUT */ EventProviders []MSFT_MTEventProvider,
	/* OUT */ WindowsEventChannels []MSFT_MTEventChannel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetProvidersAndWindowsEventChannels")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
