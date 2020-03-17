// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_SAAction struct
type CIM_SAAction struct {
	CIM_PolicyAction

	//
	DoPacketLogging bool
}

// SetDoPacketLogging sets the value of DoPacketLogging for the instance
func (instance *CIM_SAAction) SetPropertyDoPacketLogging(value bool) (err error) {
	return instance.SetProperty("DoPacketLogging", value)
}

// GetDoPacketLogging gets the value of DoPacketLogging for the instance
func (instance *CIM_SAAction) GetPropertyDoPacketLogging() (value bool, err error) {
	retValue, err := instance.GetProperty("DoPacketLogging")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
