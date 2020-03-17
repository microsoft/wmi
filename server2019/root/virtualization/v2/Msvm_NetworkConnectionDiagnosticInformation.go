// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_NetworkConnectionDiagnosticInformation struct
type Msvm_NetworkConnectionDiagnosticInformation struct {
	cim.WmiInstance

	// The round trip time for the Ping request.
	RoundTripTime uint32
}

// SetRoundTripTime sets the value of RoundTripTime for the instance
func (instance *Msvm_NetworkConnectionDiagnosticInformation) SetPropertyRoundTripTime(value uint32) (err error) {
	return instance.SetProperty("RoundTripTime", value)
}

// GetRoundTripTime gets the value of RoundTripTime for the instance
func (instance *Msvm_NetworkConnectionDiagnosticInformation) GetPropertyRoundTripTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("RoundTripTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
