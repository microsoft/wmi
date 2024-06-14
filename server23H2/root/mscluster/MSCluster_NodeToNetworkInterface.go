// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_NodeToNetworkInterface struct
type MSCluster_NodeToNetworkInterface struct {
	*CIM_SystemDevice
}

func NewMSCluster_NodeToNetworkInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_NodeToNetworkInterface, err error) {
	tmp, err := NewCIM_SystemDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToNetworkInterface{
		CIM_SystemDevice: tmp,
	}
	return
}

func NewMSCluster_NodeToNetworkInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_NodeToNetworkInterface, err error) {
	tmp, err := NewCIM_SystemDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToNetworkInterface{
		CIM_SystemDevice: tmp,
	}
	return
}
