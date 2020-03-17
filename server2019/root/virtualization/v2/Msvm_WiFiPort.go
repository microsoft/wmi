// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_WiFiPort struct
type Msvm_WiFiPort struct {
	CIM_WiFiPort

	//
	IsBound bool
}

// SetIsBound sets the value of IsBound for the instance
func (instance *Msvm_WiFiPort) SetPropertyIsBound(value bool) (err error) {
	return instance.SetProperty("IsBound", value)
}

// GetIsBound gets the value of IsBound for the instance
func (instance *Msvm_WiFiPort) GetPropertyIsBound() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBound")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
