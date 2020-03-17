// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_ExternalFcPort struct
type Msvm_ExternalFcPort struct {
	CIM_FCPort

	//
	IsHyperVCapable bool

	//
	WWNN string

	//
	WWPN string
}

// SetIsHyperVCapable sets the value of IsHyperVCapable for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyIsHyperVCapable(value bool) (err error) {
	return instance.SetProperty("IsHyperVCapable", value)
}

// GetIsHyperVCapable gets the value of IsHyperVCapable for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyIsHyperVCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHyperVCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWWNN sets the value of WWNN for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyWWNN(value string) (err error) {
	return instance.SetProperty("WWNN", value)
}

// GetWWNN gets the value of WWNN for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyWWNN() (value string, err error) {
	retValue, err := instance.GetProperty("WWNN")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWWPN sets the value of WWPN for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyWWPN(value string) (err error) {
	return instance.SetProperty("WWPN", value)
}

// GetWWPN gets the value of WWPN for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyWWPN() (value string, err error) {
	retValue, err := instance.GetProperty("WWPN")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
