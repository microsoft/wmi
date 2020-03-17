// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_IEEE8023GroupPolicySetting struct
type RSOP_IEEE8023GroupPolicySetting struct {
	RSOP_PolicySetting

	//
	description string

	//
	msieee8023PolicyData string

	//
	msieee8023PolicyReserved []uint8

	//
	whenChanged uint32
}

// Setdescription sets the value of description for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", value)
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertydescription() (value string, err error) {
	retValue, err := instance.GetProperty("description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setmsieee8023PolicyData sets the value of msieee8023PolicyData for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertymsieee8023PolicyData(value string) (err error) {
	return instance.SetProperty("msieee8023PolicyData", value)
}

// Getmsieee8023PolicyData gets the value of msieee8023PolicyData for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertymsieee8023PolicyData() (value string, err error) {
	retValue, err := instance.GetProperty("msieee8023PolicyData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setmsieee8023PolicyReserved sets the value of msieee8023PolicyReserved for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertymsieee8023PolicyReserved(value []uint8) (err error) {
	return instance.SetProperty("msieee8023PolicyReserved", value)
}

// Getmsieee8023PolicyReserved gets the value of msieee8023PolicyReserved for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertymsieee8023PolicyReserved() (value []uint8, err error) {
	retValue, err := instance.GetProperty("msieee8023PolicyReserved")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetwhenChanged sets the value of whenChanged for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) SetPropertywhenChanged(value uint32) (err error) {
	return instance.SetProperty("whenChanged", value)
}

// GetwhenChanged gets the value of whenChanged for the instance
func (instance *RSOP_IEEE8023GroupPolicySetting) GetPropertywhenChanged() (value uint32, err error) {
	retValue, err := instance.GetProperty("whenChanged")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
