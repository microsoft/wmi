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

// HNet_ConnectionAutoconfig struct
type HNet_ConnectionAutoconfig struct {
	cim.WmiInstance

	//
	Connection HNet_Connection
}

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionAutoconfig) SetPropertyConnection(value HNet_Connection) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionAutoconfig) GetPropertyConnection() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}
