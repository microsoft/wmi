// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_UpdateServices_Services struct
type ServerComponent_UpdateServices_Services struct {
	*MSFT_ServerManagerServerComponentDescriptor

	//
	ContentDirectory string

	//
	ContentLocal bool
}

func NewServerComponent_UpdateServices_ServicesEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_UpdateServices_Services, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_UpdateServices_Services{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_UpdateServices_ServicesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_UpdateServices_Services, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_UpdateServices_Services{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

// SetContentDirectory sets the value of ContentDirectory for the instance
func (instance *ServerComponent_UpdateServices_Services) SetPropertyContentDirectory(value string) (err error) {
	return instance.SetProperty("ContentDirectory", value)
}

// GetContentDirectory gets the value of ContentDirectory for the instance
func (instance *ServerComponent_UpdateServices_Services) GetPropertyContentDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ContentDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentLocal sets the value of ContentLocal for the instance
func (instance *ServerComponent_UpdateServices_Services) SetPropertyContentLocal(value bool) (err error) {
	return instance.SetProperty("ContentLocal", value)
}

// GetContentLocal gets the value of ContentLocal for the instance
func (instance *ServerComponent_UpdateServices_Services) GetPropertyContentLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("ContentLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
