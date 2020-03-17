// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

// ServerComponent_UpdateServices_Services struct
type ServerComponent_UpdateServices_Services struct {
	MSFT_ServerManagerServerComponentDescriptor

	//
	ContentDirectory string

	//
	ContentLocal bool
}

// SetContentDirectory sets the value of ContentDirectory for the instance
func (instance *ServerComponent_UpdateServices_Services) SetPropertyContentDirectory(value string) (err error) {
	return instance.SetProperty("ContentDirectory", value)
}

// GetContentDirectory gets the value of ContentDirectory for the instance
func (instance *ServerComponent_UpdateServices_Services) GetPropertyContentDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ContentDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentLocal sets the value of ContentLocal for the instance
func (instance *ServerComponent_UpdateServices_Services) SetPropertyContentLocal(value bool) (err error) {
	return instance.SetProperty("ContentLocal", value)
}

// GetContentLocal gets the value of ContentLocal for the instance
func (instance *ServerComponent_UpdateServices_Services) GetPropertyContentLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("ContentLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
