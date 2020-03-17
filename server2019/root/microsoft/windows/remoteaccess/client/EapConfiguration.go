// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// EapConfiguration struct
type EapConfiguration struct {
	cim.WmiInstance

	//
	EapConfigXmlStream string
}

// SetEapConfigXmlStream sets the value of EapConfigXmlStream for the instance
func (instance *EapConfiguration) SetPropertyEapConfigXmlStream(value string) (err error) {
	return instance.SetProperty("EapConfigXmlStream", value)
}

// GetEapConfigXmlStream gets the value of EapConfigXmlStream for the instance
func (instance *EapConfiguration) GetPropertyEapConfigXmlStream() (value string, err error) {
	retValue, err := instance.GetProperty("EapConfigXmlStream")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
