// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Account struct
type Win32_Account struct {
	CIM_LogicalElement

	//
	Domain string

	//
	LocalAccount bool

	//
	SID string

	//
	SIDType uint8
}

// SetDomain sets the value of Domain for the instance
func (instance *Win32_Account) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *Win32_Account) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalAccount sets the value of LocalAccount for the instance
func (instance *Win32_Account) SetPropertyLocalAccount(value bool) (err error) {
	return instance.SetProperty("LocalAccount", value)
}

// GetLocalAccount gets the value of LocalAccount for the instance
func (instance *Win32_Account) GetPropertyLocalAccount() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalAccount")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSID sets the value of SID for the instance
func (instance *Win32_Account) SetPropertySID(value string) (err error) {
	return instance.SetProperty("SID", value)
}

// GetSID gets the value of SID for the instance
func (instance *Win32_Account) GetPropertySID() (value string, err error) {
	retValue, err := instance.GetProperty("SID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSIDType sets the value of SIDType for the instance
func (instance *Win32_Account) SetPropertySIDType(value uint8) (err error) {
	return instance.SetProperty("SIDType", value)
}

// GetSIDType gets the value of SIDType for the instance
func (instance *Win32_Account) GetPropertySIDType() (value uint8, err error) {
	retValue, err := instance.GetProperty("SIDType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
