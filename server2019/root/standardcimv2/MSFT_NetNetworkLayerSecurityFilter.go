// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetNetworkLayerSecurityFilter struct
type MSFT_NetNetworkLayerSecurityFilter struct {
	CIM_FilterEntryBase

	//
	Authentication uint16

	//
	Encryption uint16

	//
	LocalUsers string

	//
	OverrideBlockRules bool

	//
	RemoteMachines string

	//
	RemoteUsers string
}

// SetAuthentication sets the value of Authentication for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyAuthentication(value uint16) (err error) {
	return instance.SetProperty("Authentication", value)
}

// GetAuthentication gets the value of Authentication for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyAuthentication() (value uint16, err error) {
	retValue, err := instance.GetProperty("Authentication")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryption sets the value of Encryption for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyEncryption(value uint16) (err error) {
	return instance.SetProperty("Encryption", value)
}

// GetEncryption gets the value of Encryption for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyEncryption() (value uint16, err error) {
	retValue, err := instance.GetProperty("Encryption")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalUsers sets the value of LocalUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyLocalUsers(value string) (err error) {
	return instance.SetProperty("LocalUsers", value)
}

// GetLocalUsers gets the value of LocalUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyLocalUsers() (value string, err error) {
	retValue, err := instance.GetProperty("LocalUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOverrideBlockRules sets the value of OverrideBlockRules for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyOverrideBlockRules(value bool) (err error) {
	return instance.SetProperty("OverrideBlockRules", value)
}

// GetOverrideBlockRules gets the value of OverrideBlockRules for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyOverrideBlockRules() (value bool, err error) {
	retValue, err := instance.GetProperty("OverrideBlockRules")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteMachines sets the value of RemoteMachines for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyRemoteMachines(value string) (err error) {
	return instance.SetProperty("RemoteMachines", value)
}

// GetRemoteMachines gets the value of RemoteMachines for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyRemoteMachines() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachines")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteUsers sets the value of RemoteUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) SetPropertyRemoteUsers(value string) (err error) {
	return instance.SetProperty("RemoteUsers", value)
}

// GetRemoteUsers gets the value of RemoteUsers for the instance
func (instance *MSFT_NetNetworkLayerSecurityFilter) GetPropertyRemoteUsers() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
