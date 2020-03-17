// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_BridgeMember struct
type HNet_BridgeMember struct {
	cim.WmiInstance

	//
	Bridge HNet_Connection

	//
	Member HNet_Connection
}

// SetBridge sets the value of Bridge for the instance
func (instance *HNet_BridgeMember) SetPropertyBridge(value HNet_Connection) (err error) {
	return instance.SetProperty("Bridge", value)
}

// GetBridge gets the value of Bridge for the instance
func (instance *HNet_BridgeMember) GetPropertyBridge() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Bridge")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMember sets the value of Member for the instance
func (instance *HNet_BridgeMember) SetPropertyMember(value HNet_Connection) (err error) {
	return instance.SetProperty("Member", value)
}

// GetMember gets the value of Member for the instance
func (instance *HNet_BridgeMember) GetPropertyMember() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Member")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}
