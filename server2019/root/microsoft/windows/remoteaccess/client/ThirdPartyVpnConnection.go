// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

// ThirdPartyVpnConnection struct
type ThirdPartyVpnConnection struct {
	VpnCommonConfig

	//
	CustomConfiguration string

	//
	PlugInApplicationID string

	//
	VpnConfigurationXml string
}

// SetCustomConfiguration sets the value of CustomConfiguration for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyCustomConfiguration(value string) (err error) {
	return instance.SetProperty("CustomConfiguration", value)
}

// GetCustomConfiguration gets the value of CustomConfiguration for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyCustomConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("CustomConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlugInApplicationID sets the value of PlugInApplicationID for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyPlugInApplicationID(value string) (err error) {
	return instance.SetProperty("PlugInApplicationID", value)
}

// GetPlugInApplicationID gets the value of PlugInApplicationID for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyPlugInApplicationID() (value string, err error) {
	retValue, err := instance.GetProperty("PlugInApplicationID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnConfigurationXml sets the value of VpnConfigurationXml for the instance
func (instance *ThirdPartyVpnConnection) SetPropertyVpnConfigurationXml(value string) (err error) {
	return instance.SetProperty("VpnConfigurationXml", value)
}

// GetVpnConfigurationXml gets the value of VpnConfigurationXml for the instance
func (instance *ThirdPartyVpnConnection) GetPropertyVpnConfigurationXml() (value string, err error) {
	retValue, err := instance.GetProperty("VpnConfigurationXml")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
