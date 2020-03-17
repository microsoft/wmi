// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_IKEAction struct
type CIM_IKEAction struct {
	CIM_SANegotiationAction

	//
	AggressiveModeGroupID uint16

	//
	ExchangeMode uint16

	//
	UseIKEIdentityType uint16

	//
	VendorID string
}

// SetAggressiveModeGroupID sets the value of AggressiveModeGroupID for the instance
func (instance *CIM_IKEAction) SetPropertyAggressiveModeGroupID(value uint16) (err error) {
	return instance.SetProperty("AggressiveModeGroupID", value)
}

// GetAggressiveModeGroupID gets the value of AggressiveModeGroupID for the instance
func (instance *CIM_IKEAction) GetPropertyAggressiveModeGroupID() (value uint16, err error) {
	retValue, err := instance.GetProperty("AggressiveModeGroupID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExchangeMode sets the value of ExchangeMode for the instance
func (instance *CIM_IKEAction) SetPropertyExchangeMode(value uint16) (err error) {
	return instance.SetProperty("ExchangeMode", value)
}

// GetExchangeMode gets the value of ExchangeMode for the instance
func (instance *CIM_IKEAction) GetPropertyExchangeMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExchangeMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseIKEIdentityType sets the value of UseIKEIdentityType for the instance
func (instance *CIM_IKEAction) SetPropertyUseIKEIdentityType(value uint16) (err error) {
	return instance.SetProperty("UseIKEIdentityType", value)
}

// GetUseIKEIdentityType gets the value of UseIKEIdentityType for the instance
func (instance *CIM_IKEAction) GetPropertyUseIKEIdentityType() (value uint16, err error) {
	retValue, err := instance.GetProperty("UseIKEIdentityType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_IKEAction) SetPropertyVendorID(value string) (err error) {
	return instance.SetProperty("VendorID", value)
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_IKEAction) GetPropertyVendorID() (value string, err error) {
	retValue, err := instance.GetProperty("VendorID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
