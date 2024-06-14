// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_Property struct
type MSCluster_Property struct {
	*cim.WmiInstance
}

func NewMSCluster_PropertyEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Property, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_Property{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_PropertyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Property, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Property{
		WmiInstance: tmp,
	}
	return
}
