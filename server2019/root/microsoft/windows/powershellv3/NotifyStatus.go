// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NotifyStatus struct
type __NotifyStatus struct {
	cim.WmiInstance

	//
	StatusCode uint32
}

// SetStatusCode sets the value of StatusCode for the instance
func (instance *__NotifyStatus) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", value)
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *__NotifyStatus) GetPropertyStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
