// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_SARule struct
type CIM_SARule struct {
	CIM_PolicyRule

	//
	LimitNegotiation uint16
}

// SetLimitNegotiation sets the value of LimitNegotiation for the instance
func (instance *CIM_SARule) SetPropertyLimitNegotiation(value uint16) (err error) {
	return instance.SetProperty("LimitNegotiation", value)
}

// GetLimitNegotiation gets the value of LimitNegotiation for the instance
func (instance *CIM_SARule) GetPropertyLimitNegotiation() (value uint16, err error) {
	retValue, err := instance.GetProperty("LimitNegotiation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
