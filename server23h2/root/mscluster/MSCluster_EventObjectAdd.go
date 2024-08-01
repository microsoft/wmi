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

// MSCluster_EventObjectAdd struct
type MSCluster_EventObjectAdd struct {
	*MSCluster_Event
}

func NewMSCluster_EventObjectAddEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventObjectAdd, err error) {
	tmp, err := NewMSCluster_EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventObjectAdd{
		MSCluster_Event: tmp,
	}
	return
}

func NewMSCluster_EventObjectAddEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_EventObjectAdd, err error) {
	tmp, err := NewMSCluster_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventObjectAdd{
		MSCluster_Event: tmp,
	}
	return
}
