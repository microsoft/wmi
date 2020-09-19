// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClusteringService struct
type CIM_ClusteringService struct {
	*CIM_Service
}

func NewCIM_ClusteringServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_ClusteringService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClusteringService{
		CIM_Service: tmp,
	}
	return
}

func NewCIM_ClusteringServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClusteringService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClusteringService{
		CIM_Service: tmp,
	}
	return
}
