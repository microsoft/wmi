// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_IEESCSecurityZoneSettings struct
type RSOP_IEESCSecurityZoneSettings struct {
	RSOP_IESecurityZoneSettings

	//
	EscEnabled bool
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", value)
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) GetPropertyEscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EscEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
