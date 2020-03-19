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

// ServerComponent_UpdateServices_Database struct
type ServerComponent_UpdateServices_Database struct {
	*MSFT_ServerManagerServerComponentDescriptor

	//
	InstanceName string
}

func NewServerComponent_UpdateServices_DatabaseEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_UpdateServices_Database, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_UpdateServices_Database{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_UpdateServices_DatabaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_UpdateServices_Database, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_UpdateServices_Database{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ServerComponent_UpdateServices_Database) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", value)
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ServerComponent_UpdateServices_Database) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
