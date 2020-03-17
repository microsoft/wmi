// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetLbfoTeamNic struct
type MSFT_NetLbfoTeamNic struct {
	MSFT_NetImPlatAdapter

	// 401
	Default bool

	// 400
	Primary bool

	// 399
	VlanID uint32
}

// SetDefault sets the value of Default for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyDefault(value bool) (err error) {
	return instance.SetProperty("Default", value)
}

// GetDefault gets the value of Default for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("Default")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimary sets the value of Primary for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyPrimary(value bool) (err error) {
	return instance.SetProperty("Primary", value)
}

// GetPrimary gets the value of Primary for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyPrimary() (value bool, err error) {
	retValue, err := instance.GetProperty("Primary")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyVlanID(value uint32) (err error) {
	return instance.SetProperty("VlanID", value)
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyVlanID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VlanID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
