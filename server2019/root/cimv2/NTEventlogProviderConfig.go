// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// NTEventlogProviderConfig struct
type NTEventlogProviderConfig struct {
	cim.WmiInstance

	//
	LastBootUpTime string
}

// SetLastBootUpTime sets the value of LastBootUpTime for the instance
func (instance *NTEventlogProviderConfig) SetPropertyLastBootUpTime(value string) (err error) {
	return instance.SetProperty("LastBootUpTime", value)
}

// GetLastBootUpTime gets the value of LastBootUpTime for the instance
func (instance *NTEventlogProviderConfig) GetPropertyLastBootUpTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastBootUpTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
