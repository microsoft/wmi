// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_WiFiEndpoint struct
type Msvm_WiFiEndpoint struct {
	CIM_WiFiEndpoint

	//
	Connected bool
}

// SetConnected sets the value of Connected for the instance
func (instance *Msvm_WiFiEndpoint) SetPropertyConnected(value bool) (err error) {
	return instance.SetProperty("Connected", value)
}

// GetConnected gets the value of Connected for the instance
func (instance *Msvm_WiFiEndpoint) GetPropertyConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("Connected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
