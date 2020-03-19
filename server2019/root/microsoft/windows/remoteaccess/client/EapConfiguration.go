// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// EapConfiguration struct
type EapConfiguration struct {
	*cim.WmiInstance

	//
	EapConfigXmlStream string
}

func NewEapConfigurationEx1(instance *cim.WmiInstance) (newInstance *EapConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &EapConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewEapConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *EapConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &EapConfiguration{
		WmiInstance: tmp,
	}
	return
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
